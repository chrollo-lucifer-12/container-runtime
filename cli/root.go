package cli

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "monster",
		SilenceUsage: true,
	}

	cmd.AddCommand(startCmd(), createCmd(), deleteCmd(), killCmd(), stateCmd())

	return cmd
}
