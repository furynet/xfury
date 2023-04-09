package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/xfury/utils"
	"github.com/furynet/xfury/x/dvm/types"
)

// SetProposalStats sets proposal statistics in the store
func (k Keeper) SetProposalStats(ctx sdk.Context, stats types.ProposalStats) {
	store := k.getProposalStatStore(ctx)
	b := k.cdc.MustMarshal(&stats)
	store.Set(utils.StrBytes("0"), b)
}

// GetProposalStats returns proposal stats
func (k Keeper) GetProposalStats(ctx sdk.Context) (val types.ProposalStats) {
	store := k.getProposalStatStore(ctx)

	b := store.Get(utils.StrBytes("0"))
	if b == nil {
		return val
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}
