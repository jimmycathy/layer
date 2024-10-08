package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/tellor-io/layer/x/registry/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDataSpec(goCtx context.Context, req *types.MsgUpdateDataSpec) (*types.MsgUpdateDataSpecResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	// normalize query type
	req.QueryType = strings.ToLower(req.QueryType)
	// check if the query type exists
	querytypeExists, err := k.Keeper.HasSpec(ctx, req.QueryType)
	if err != nil {
		return nil, err
	}
	if !querytypeExists {
		return nil, errorsmod.Wrapf(types.ErrInvalidSpec, "data spec not registered for query type: %s", req.QueryType)
	}
	if err := k.Keeper.SetDataSpec(ctx, req.QueryType, req.Spec); err != nil {
		return nil, err
	}

	if err := k.Keeper.Hooks().AfterDataSpecUpdated(ctx, req.QueryType, req.Spec); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"data_spec_updated",
			sdk.NewAttribute("query_type", req.QueryType),
			sdk.NewAttribute("document_hash_id", req.Spec.DocumentHash),
			sdk.NewAttribute("aggregate_method", req.Spec.AggregationMethod),
			sdk.NewAttribute("response_value_type", req.Spec.ResponseValueType),
			sdk.NewAttribute("report_buffer_window", fmt.Sprintf("%d", req.Spec.ReportBlockWindow)),
		),
	})
	return &types.MsgUpdateDataSpecResponse{}, nil
}
