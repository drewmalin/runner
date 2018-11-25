package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/drewmalin/runner/internal/cmdcontext"
)

const command = "version"

// NewCmdVersion builds the 'version' subcommand of the root 'runner' command.
func NewCmdVersion(context *cmdcontext.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: command,
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := buildVersionString(context)
			if err != nil {
				return err
			}

			fmt.Fprintln(context.VerboseOut, version)
			return nil
		},
	}
	return cmd
}