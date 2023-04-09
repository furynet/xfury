package cli_test

import (
	"strings"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/furynet/fury/testutil/network"
	"github.com/furynet/fury/x/mint/client/cli"
	"github.com/stretchr/testify/require"
)

func TestGetQueryCmd(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	for _, tc := range []struct {
		desc string
		args []string
		err  error
		code uint32
	}{
		{
			desc: "valid",
			args: []string{},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{}
			args = append(args, tc.args...)
			res, err := clitestutil.ExecTestCLICmd(ctx, cli.GetQueryCmd(""), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}

			require.True(t, strings.HasPrefix(string(res.Bytes()), "Querying commands for the mint module"))
		})
	}
}
