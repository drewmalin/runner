package version

import (
	"github.com/drewmalin/runner/internal/cmdcontext"
	"github.com/drewmalin/runner/internal/config"
)

// buildVersionString returns the full output for the 'version' command. The result of
// this function may be immediately printed to standard out.
func buildVersionString(context *cmdcontext.Context) (string, error) {
	return config.GetRunnerVersion(), nil
}