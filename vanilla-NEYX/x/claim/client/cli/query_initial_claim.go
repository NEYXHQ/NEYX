package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/ignite/modules/x/claim/types"
)

func CmdShowInitialClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-initial-claim",
		Short: "shows information about initial claim",
		Long:  "shows if initial claim is enabled and what is the mission ID completed by initial claim",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetInitialClaimRequest{}

			res, err := queryClient.InitialClaim(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
