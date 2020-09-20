package config

var (
	// Executable is the name of the exec
	Executable = "fluoride"
	// Version is the current version
	Version = "1.0.0"
)

func init() {
	Executable = GetConfig().Application.Name
	Version = GetConfig().Application.Version
}
