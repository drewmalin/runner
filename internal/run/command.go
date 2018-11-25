package run

import (
	"github.com/spf13/cobra"

	"github.com/drewmalin/runner/internal/cmdcontext"
)

const command = "run"

// NewCmdVersion builds the 'version' subcommand of the root 'runner' command.
func NewCmdRun(context *cmdcontext.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: command,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := run(context)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}