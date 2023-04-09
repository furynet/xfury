package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simappUtil "github.com/furynet/xfury/testutil/simapp"
	"github.com/furynet/xfury/x/dvm/keeper"
	"github.com/stretchr/testify/require"
)

func setupKeeperAndApp(t testing.TB) (*simappUtil.TestApp, *keeper.KeeperTest, sdk.Context) {
	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	return tApp, &tApp.DVMKeeper, ctx
}

func setupKeeper(t testing.TB) (*keeper.KeeperTest, sdk.Context) {
	_, k, ctx := setupKeeperAndApp(t)

	return k, ctx
}

func TestLogger(t *testing.T) {
	k, ctx := setupKeeper(t)
	logger := k.Logger(ctx)
	require.True(t, logger != nil)
}
