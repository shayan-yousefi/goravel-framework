package support

const Version string = "v1.12.6"

const (
	EnvRuntime = "runtime"
	EnvArtisan = "artisan"
	EnvTest    = "test"
)

var (
	Env      = EnvRuntime
	RootPath string
)
