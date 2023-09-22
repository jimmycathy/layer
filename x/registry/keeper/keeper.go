package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/tellor-io/layer/x/registry/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetGenesisSpec(ctx sdk.Context) {
	var dataSpec types.DataSpec
	dataSpec.DocumentHash = ""
	dataSpec.ValueType = "uint256"
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpecRegistryKey))
	store.Set([]byte("SpotPrice"), k.cdc.MustMarshal(&dataSpec))

}

func (k Keeper) GetGenesisSpec(ctx sdk.Context) types.DataSpec {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpecRegistryKey))
	spec := store.Get([]byte("SpotPrice"))
	var dataSpec types.DataSpec
	k.cdc.Unmarshal(spec, &dataSpec)
	return dataSpec
}

func (k Keeper) SetGenesisQuery(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QueryRegistryKey))

	encodedData, _ := EncodeArguments([]string{"string", "string"}, []string{"eth", "usd"})

	queryData, _ := EncodeArguments([]string{"string", "bytes"}, []string{"SpotPrice", string(encodedData)})

	queryId := crypto.Keccak256(queryData)
	store.Set(queryId, queryData)

}

func (k Keeper) GetGenesisQuery(ctx sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SpecRegistryKey))
	queryId, _ := k.QueryData(ctx, "83A7F3D48786AC2667503A61E8C415438ED2922EB86A2906E4EE66D9A2CE4992")
	queryData := store.Get(queryId)
	return bytes.HexBytes(queryData).String()
}
