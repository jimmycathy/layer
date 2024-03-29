package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/utils"
	"github.com/tellor-io/layer/x/oracle/types"
)

func (k msgServer) CommitReport(goCtx context.Context, msg *types.MsgCommitReport) (*types.MsgCommitReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reporterAddr, err := msg.GetSignerAndValidateMsg()
	if err != nil {
		return nil, err
	}

	// get reporter
	reporter, err := k.reporterKeeper.Reporter(ctx, reporterAddr)
	if err != nil {
		return nil, err
	}
	if reporter.Jailed {
		return nil, errorsmod.Wrapf(types.ErrReporterJailed, "reporter %s is in jail", reporterAddr)
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	if reporter.TotalTokens.LT(params.MinStakeAmount) {
		return nil, errorsmod.Wrapf(types.ErrNotEnoughStake, "reporter has %s, required amount is %s", reporter.TotalTokens, params.MinStakeAmount)
	}
	// normalize by removing 0x prefix and lowercasing
	msg.QueryData = utils.Remove0xPrefix(msg.QueryData)

	// get query id bytes hash from query data
	queryId, err := utils.QueryIDFromDataString(msg.QueryData)
	if err != nil {
		return nil, types.ErrInvalidQueryData.Wrapf("invalid query data: %s", err)
	}
	// get query info by query id
	query, err := k.Keeper.Query.Get(ctx, queryId)
	if err != nil {
		// if no query it means its not a cyclelist query and doesn't have tips (cyclelist queries are initialized in genesis)
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrNoTipsNotInCycle.Wrapf("query not part of cyclelist")
		}
		return nil, err
	}
	// get current query in cycle
	cycleQuery, err := k.Keeper.GetCurrentQueryInCycleList(ctx)
	if err != nil {
		return nil, err
	}
	// bool to check if query is in cycle
	incycle := msg.QueryData == cycleQuery

	if !incycle {
		k.Logger(ctx).Info("query not in cycle")
	}

	k.Logger(ctx).Info("Expiration", "exp", query.Expiration)
	k.Logger(ctx).Info("BlockTime", "blocktime", ctx.BlockTime())
	k.Logger(ctx).Info("offset", "offset", offset)

	if query.Expiration.Before(ctx.BlockTime()) {
		k.Logger(ctx).Info("query expired")
	}

	if query.Amount.IsZero() {
		k.Logger(ctx).Info("query does not have tips and is not in cycle")
	}

	if query.Amount.IsZero() && query.Expiration.Before(ctx.BlockTime()) && !incycle {
		return nil, types.ErrNoTipsNotInCycle.Wrapf("query does not have tips and is not in cycle")
	}

	if query.Amount.GT(math.ZeroInt()) && query.Expiration.Before(ctx.BlockTime()) && !incycle {
		return nil, errors.New("query's tip is expired and is not in cycle")
	}

	// if tip is zero and expired, move query forward only if in cycle
	// if tip amount is zero and query timeframe is expired, it means one of two things
	// the tip has been paid out because the query has expired and there were revealed reports
	// or the query was in cycle and expired (either revealed or not)
	// in either case move query forward by incrementing id and setting expiration
	if query.Amount.IsZero() && query.Expiration.Before(ctx.BlockTime()) && incycle {
		nextId, err := k.Keeper.QuerySequnecer.Next(ctx)
		if err != nil {
			return nil, err
		}
		query.Id = nextId
		// reset query fields when generating next id
		query.HasRevealedReports = false
		query.Expiration = ctx.BlockTime().Add(offset)
		err = k.Query.Set(ctx, queryId, query)
		if err != nil {
			return nil, err
		}
	}

	// if there is tip but window expired, only incycle can extend the window, otherwise requires tip vi msgTip tx
	// if tip amount is greater than zero and query timeframe is expired, it means that the query didn't have any revealed reports
	// and the tip is still there and so the time can be extended only if the query is in cycle or via a tip transaction
	// maintains the same id until the query is paid out
	if query.Amount.GT(math.ZeroInt()) && query.Expiration.Before(ctx.BlockTime()) && incycle {
		query.Expiration = ctx.BlockTime().Add(query.RegistrySpecTimeframe)
		err = k.Query.Set(ctx, queryId, query)
		if err != nil {
			return nil, err
		}
	}

	// if tip is zero and not expired, this could only mean that the query is still accepting submissions
	if query.Amount.IsZero() && ctx.BlockTime().Before(query.Expiration) {
		incycle = true
	}

	commit := types.Commit{
		Reporter: msg.Creator,
		QueryId:  queryId,
		Hash:     msg.Hash,
		Incycle:  incycle,
	}
	err = k.Commits.Set(ctx, collections.Join(reporterAddr.Bytes(), query.Id), commit)
	if err != nil {
		return nil, err
	}

	return &types.MsgCommitReportResponse{}, nil
}
