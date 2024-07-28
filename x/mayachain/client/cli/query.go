package cli

import (
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

type ver struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildTime string `json:"build_time"`
}

func (v ver) String() string {
	return v.Version
}

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the BASEChain module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdGetVersion())
	cmd.AddCommand(GetCmdGetNORelay())
	return cmd
}

// GetCmdGetVersion queries current version
func GetCmdGetVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Gets the BASEChain version and build information",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			clientCtx.OutputFormat = "json"

			out := ver{
				Version:   constants.SWVersion.String(),
				GitCommit: constants.GitCommit,
				BuildTime: constants.BuildTime,
			}
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdGetNORelay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discord-relay",
		Short: "Relays a message from a node operator to discord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			msg := common.NewNodeRelay(args[0], args[1])

			if err = msg.Prepare(); err != nil {
				log.Fatalln(err)
			}

			result, err := msg.Broadcast()
			if err != nil {
				log.Fatalln(err)
			}

			return clientCtx.PrintObjectLegacy(result)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
