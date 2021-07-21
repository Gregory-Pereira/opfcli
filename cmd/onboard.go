package cmd

import (
	"fmt"
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdOnboard(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use: "opfcli onboard file-path",
		Long: "The onboarding command creates all the necesary components to onboard a team: Namespaces, groups, role bindings, and allows them to be accessed on the cluster",
		Short: `Onboard a new project, group, and namespace(s) into Operate First.`,
		Args: cobra.ExactArgs(1),
		SilenceUsage: true,
		SilenceErrors true,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDisplayName, err := cmd.Flags().GetString("display-name")
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			disableLimitrange, err := cmd.Flags().GetBool("no-limitrange")
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			return opfapi.Onboard(args[0], projectDisplayName, disableLimitrange)
		},
	}
	cmd.Flags().StringP("display-name", "d", "", "Short team description for easy identification of project")
	cmd.Flags().BoolP("no-limitrange", "n", false, "Do not set a limitrange on this project")
	return cmd
}
