package config

type configs struct {}

var config configs

// Below vars are exposed to the build-layer (Makefile) so that they be overridden at build time.
var (
	RunnerVersion = "unknown"
)

func Init() {
	config = configs{}
}

// GetRunnerVersion returns the version of the 'runner' command. The value here is set at build-
// time.
func GetRunnerVersion() string {
	return RunnerVersion
}