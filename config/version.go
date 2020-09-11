package config

import (
	"github.com/spf13/viper"
)

var (
	// Executable is the name of the exec
	Executable = "fluoride"
	// Version is the current version
	Version = "1.0.0"
)

func init() {
	Version = viper.GetString("application.version")
}
