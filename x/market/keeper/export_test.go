package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/fury/x/market/types"
)

// KeeperTest is a wrapper object for the keeper, It is being used
// to export unexported methods of the keeper
type KeeperTest = Keeper

func (k KeeperTest) ValidateMarketResolution(resolutionPayload types.MarketResolutionTicketPayload) error {
	return resolutionPayload.Validate()
}

func (k KeeperTest) ValidateMarketAdd(ctx sdk.Context, addPayload types.MarketAddTicketPayload) error {
	params := k.GetParams(ctx)
	return addPayload.Validate(ctx, &params)
}

func (k KeeperTest) ValidateMarketUpdate(ctx sdk.Context, updatePayload types.MarketUpdateTicketPayload, currentMarket types.Market) error {
	params := k.GetParams(ctx)
	return updatePayload.Validate(ctx, &params)
}
