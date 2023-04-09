package simulation

import (
	//#nosec
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/furynet/xfury/x/market/keeper"
	"github.com/furynet/xfury/x/market/types"
)

// SimulateMsgAddMarket simulates the add market flow
func SimulateMsgAddMarket(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddMarket{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddMarket simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddMarket simulation not implemented"), nil, nil
	}
}
