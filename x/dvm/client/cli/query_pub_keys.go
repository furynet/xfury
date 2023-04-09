package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/furynet/fury/x/dvm/types"
	"github.com/spf13/cobra"
)

// CmdPubKeysList returns a command object instance for querying the public keys
func CmdPubKeysList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pubkeys",
		Short: "Query public keys list",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPubKeysRequest{}

			res, err := queryClient.PubKeys(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
