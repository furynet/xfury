package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/fury/x/market/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

// GetDefaultBetConstraints get bet constraint values of the bet constraints
func (k Keeper) GetDefaultBetConstraints(ctx sdk.Context) (params *types.MarketBetConstraints) {
	p := k.GetParams(ctx)
	return p.NewMarketBetConstraints(p.MinBetAmount, p.MinBetFee)
}
