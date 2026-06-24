package cli

import (
	"github.com/container-runtime/operations"
	"github.com/spf13/cobra"
)

func reexecCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "reexec [flags] CONTAINER_ID",
		Args:   cobra.ExactArgs(1),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			containerID := args[0]

			return operations.Reexec(&operations.ReexecOpts{
				ID: containerID,
			})
		},
	}

	return cmd
}
