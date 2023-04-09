package dvm_test

import (
	"testing"

	"github.com/furynet/xfury/testutil/nullify"
	simappUtil "github.com/furynet/xfury/testutil/simapp"
	"github.com/furynet/xfury/x/dvm"
	"github.com/furynet/xfury/x/dvm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		KeyVault: types.KeyVault{
			PublicKeys: []string{"Key1"},
		},
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	dvm.InitGenesis(ctx, tApp.DVMKeeper, genesisState)
	got := dvm.ExportGenesis(ctx, tApp.DVMKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
