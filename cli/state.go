package cli

import (
	"fmt"

	"github.com/container-runtime/operations"
	"github.com/spf13/cobra"
)

func stateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:  "state [flags] CONTAINER_ID",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			containerID := args[0]

			state, err := operations.State(&operations.StateOpts{
				ID: containerID,
			})

			if err != nil {
				return err
			}

			fmt.Println(state)

			return nil
		},
	}

	return cmd
}
