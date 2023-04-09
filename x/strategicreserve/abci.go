package strategicreserve

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/fury/x/strategicreserve/keeper"
)

// EndBlocker settles the active deposits of resolved order books
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.BatchOrderBookSettlements(ctx)
	if err != nil {
		panic(fmt.Sprintf("end block no %d failed : %s", ctx.BlockHeight(), err.Error()))
	}
}
