package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdCreateOnboard(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		use: "",
		long: "",
		short: ``,
		Args: cobra.ExactArgs(5),
		SilenceUsage: true,
		SilenceErrors true,
		RunE: func(cmd *cobra.Command, args []string) error {
		  teamName := args[0]
		  projectNames := args[1]
	    projectDescription := args[2]
	    projectQuotaTier := args[3]

	    targetCluster, err := cmd.Flags().GetString("cluster")
			if err != nil {
				return err
			}

			users, err := cmd.Flags().GetString("users")
			if err != nil {
				return err
			}

			customQuota, err := cmd.Flags().GetString("custom-quota")
			if err != nil {
				return err
			}

			gpgKeys, err := cmd.Flags().GetString("gpg-keys")
			if err != nil {
				return err
			}

			return opfapi.Onboard(
				teamName, projectNames, projectDescription,
				projectQuota,
				disableLimitrange,
			)
		}
	}
	cmd.Flags().StringP("cluster", "tc", "", "Target cluster")
	cmd.Flags().StringP("users", "u", "", "Users")
	cmd.Flags().StringP("custom-quota", "cq", "", "Custom quota")
	cmd.Flags().StringP("gpg-keys", "k", "", "GPG keys")
	return cmd;
}