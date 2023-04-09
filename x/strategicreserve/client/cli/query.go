package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/furynet/fury/x/strategicreserve/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	orderBookQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the strategicreserve module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	orderBookQueryCmd.AddCommand(
		CmdQueryParams(),
		GetCmdQueryOrderBooks(),
		GetCmdQueryOrderBook(),
		GetCmdQueryOrderBookParticipations(),
		GetCmdQueryOrderBookParticipation(),
		GetCmdQueryOrderBookExposures(),
		GetCmdQueryOrderBookExposure(),
		GetCmdQueryParticipationExposures(),
		GetCmdQueryParticipationExposure(),
		GetCmdQueryHistoricalParticipationExposures(),
		GetCmdQueryParticipationBets(),
	)

	return orderBookQueryCmd
}
