package cmd

import (
	"net"
	"net/http"
	"strings"

	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile string
	logger     *zap.SugaredLogger
	rootCmd    = &cobra.Command{
		Short: "Fluoride is a robust icon pack management tool",
		Long:  `A robust and complete icon pack management platform, built to make the lives of icon pack designers easier.`,
		Use:   strings.ToLower(config.Executable),
	}
)

func initConfiguration() {
	// Load the configuration
	config.Configure(configFile)
}

func initLogger() {

	// Build the logger
	globalLogger := config.ConfigureLogger()

	zap.ReplaceGlobals(globalLogger)
	logger = globalLogger.Sugar().With("package", "cmd")

}

func initProfiler() {
	conf := config.GetConfig().Profiler
	if conf.Enabled {
		hostPort := net.JoinHostPort(conf.Host, conf.Port)
		go http.ListenAndServe(hostPort, nil)
		logger.Infof("Profiler enabled on http://%s", hostPort)
	}
}

func init() {
	cobra.OnInitialize(initConfiguration, initLogger, initProfiler)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (config.yml by default)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Add commands
	rootCmd.AddCommand(serverCmd)
}

// Execute starts the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Errorf("%s", err.Error())
	}
}
