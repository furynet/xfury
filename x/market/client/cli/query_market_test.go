package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/furynet/xfury/testutil/network"
	"github.com/furynet/xfury/testutil/nullify"
	"github.com/furynet/xfury/x/market/client/cli"
	"github.com/furynet/xfury/x/market/types"
)

func networkWithMarketObjects(t *testing.T, n int) (*network.Network, []types.Market) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		market := types.Market{
			UID:                    cast.ToString(i),
			WinnerOddsUIDs:         []string{},
			SrContributionForHouse: sdk.NewInt(2),
		}
		nullify.Fill(&market)
		state.MarketList = append(state.MarketList, market)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.MarketList
}

func TestQueryMarketCLI(t *testing.T) {
	net, objs := networkWithMarketObjects(t, 5)
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	t.Run("ShowMarket", func(t *testing.T) {
		ctx := net.Validators[0].ClientCtx
		for _, tc := range []struct {
			desc  string
			idUID string

			args []string
			err  error
			obj  types.Market
		}{
			{
				desc:  "found",
				idUID: objs[0].UID,

				args: common,
				obj:  objs[0],
			},
			{
				desc:  "not found",
				idUID: cast.ToString(100000),

				args: common,
				err:  status.Error(codes.NotFound, "not found"),
			},
		} {
			tc := tc
			t.Run(tc.desc, func(t *testing.T) {
				args := []string{
					tc.idUID,
				}
				args = append(args, tc.args...)
				out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMarket(), args)
				if tc.err != nil {
					stat, ok := status.FromError(tc.err)
					require.True(t, ok)
					require.ErrorIs(t, stat.Err(), tc.err)
				} else {
					require.NoError(t, err)
					var resp types.QueryMarketResponse
					require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
					require.NotNil(t, resp.Market)
					require.Equal(t,
						nullify.Fill(&tc.obj),
						nullify.Fill(&resp.Market),
					)
				}
			})
		}
	})
}
