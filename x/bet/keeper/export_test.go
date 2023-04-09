package keeper

import (
	"github.com/furynet/fury/x/bet/types"
	markettypes "github.com/furynet/fury/x/market/types"
)

// KeeperTest is a wrapper object for the keeper, It is being used
// to export unexported methods of the keeper
type KeeperTest = Keeper

func (k KeeperTest) ProcessBetResultAndStatus(bet *types.Bet, market markettypes.Market) error {
	return processBetResultAndStatus(bet, market)
}

func (k KeeperTest) CheckBetStatus(bet *types.Bet) error {
	return checkBetStatus(bet.Status)
}
