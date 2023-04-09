package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furynet/fury/x/strategicreserve/types"
)

// Get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GetMaxOrderBookParticipationsAllowed(ctx),
		k.GetBatchSettlementCount(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetMaxOrderBookParticipationsAllowed - Max number of book participations allowed.
func (k Keeper) GetMaxOrderBookParticipationsAllowed(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxOrderBookParticipations, &res)
	return
}

// GetBatchSettlementCount - number of deposits to be settled in end blocker.
func (k Keeper) GetBatchSettlementCount(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyBatchSettlementCount, &res)
	return
}
