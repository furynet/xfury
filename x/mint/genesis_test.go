package mint_test

import (
	"testing"

	"github.com/furynet/xfury/testutil/nullify"
	simappUtil "github.com/furynet/xfury/testutil/simapp"
	"github.com/furynet/xfury/x/mint"
	"github.com/furynet/xfury/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	mint.InitGenesis(ctx, tApp.MintKeeper, genesisState)
	got := mint.ExportGenesis(ctx, tApp.MintKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
