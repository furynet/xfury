package market_test

import (
	"testing"

	"github.com/furynet/fury/testutil/nullify"
	simappUtil "github.com/furynet/fury/testutil/simapp"
	market "github.com/furynet/fury/x/market"
	"github.com/furynet/fury/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MarketList: []types.Market{
			{
				UID: "0",
			},
			{
				UID: "1",
			},
		},
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	market.InitGenesis(ctx, tApp.MarketKeeper, genesisState)
	got := market.ExportGenesis(ctx, tApp.MarketKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MarketList, got.MarketList)
}
