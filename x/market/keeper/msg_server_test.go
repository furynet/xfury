package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simappUtil "github.com/furynet/fury/testutil/simapp"
	"github.com/furynet/fury/x/market/keeper"
	"github.com/furynet/fury/x/market/types"
)

func setupMsgServerAndKeeper(t testing.TB) (*keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	_, k, msgk, ctx, wctx := setupMsgServerAndApp(t)
	return k, msgk, ctx, wctx
}

func setupMsgServerAndApp(t testing.TB) (*simappUtil.TestApp, *keeper.KeeperTest, types.MsgServer, sdk.Context, context.Context) {
	tApp, k, ctx := setupKeeperAndApp(t)
	return tApp, k, keeper.NewMsgServerImpl(*k), ctx, sdk.WrapSDKContext(ctx)
}
