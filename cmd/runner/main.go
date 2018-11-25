// Package main serves as the entry point for the 'runner' command.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/drewmalin/runner/internal/cmdcontext"
	"github.com/drewmalin/runner/internal/config"
	"github.com/drewmalin/runner/internal/run"
	"github.com/drewmalin/runner/internal/version"
)

func main() {
	config.Init()

	cmd := newCmdRunner()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// newCmdRunner builds the full tree of commands and sub-commands. Flags and checks that should be performed
// for every sub-command will be defined here using the Cobra library's "persistent" constructs.
func newCmdRunner() *cobra.Command {
	var verbose bool
	
	cmdContext := &cmdcontext.Context{}
	runnerCommand := &cobra.Command{
		Use: "runner",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmdContext.Init(verbose)
		},
	}

	runnerCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	
	createCmdTree(runnerCommand, cmdContext)
	return runnerCommand
}

// createCmdTree builds the tree of sub-commands under the root "runner" command.
func createCmdTree(runnerCommand *cobra.Command, context *cmdcontext.Context) {
	runnerCommand.AddCommand(version.NewCmdVersion(context))
	runnerCommand.AddCommand(run.NewCmdRun(context))
}
