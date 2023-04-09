package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/xfury/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	response, err := k.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestParamsInvalidQuery(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	_, err := k.Params(wctx, nil)
	require.Error(t, err)
}
