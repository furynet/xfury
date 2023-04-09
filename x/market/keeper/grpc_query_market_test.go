package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furynet/xfury/consts"
	"github.com/furynet/xfury/testutil/nullify"
	"github.com/furynet/xfury/x/market/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMarketQuerySingle(t *testing.T) {
	k, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMarket(k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryMarketRequest
		response *types.QueryMarketResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryMarketRequest{
				Uid: msgs[0].UID,
			},
			response: &types.QueryMarketResponse{Market: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryMarketRequest{
				Uid: msgs[1].UID,
			},
			response: &types.QueryMarketResponse{Market: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryMarketRequest{
				Uid: cast.ToString(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.Market(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
