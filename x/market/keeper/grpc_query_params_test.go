package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/fury/consts"
	"github.com/furynet/fury/x/market/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	t.Run("Success", func(t *testing.T) {
		response, err := k.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
	})

	t.Run("Failed", func(t *testing.T) {
		response, err := k.Params(wctx, nil)
		require.Equal(t, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest), err)
		require.Nil(t, response)
	})
}
