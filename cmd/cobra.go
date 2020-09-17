package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/PrunedNeuron/Fluoride/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	configuration string
	logger        *zap.SugaredLogger
	pidFile       string
	rootCmd       = &cobra.Command{
		Version: config.Version,
		Use:     config.Executable,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello world.")

		},
	}
)

func initConfiguration() {
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configuration != "" {
		viper.SetConfigFile(configuration)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read configuration file!: %s ERROR: %s\n", configuration, err.Error())
			os.Exit(1)
		}
	}
}

func initLogger() {

	loggerConfiguration := zap.NewProductionConfig()

	// Log Level
	var logLevel zapcore.Level
	if err := logLevel.Set(viper.GetString("logger.level")); err != nil {
		zap.S().Fatalw("Could not determine logger.level", "error", err)
	}
	loggerConfiguration.Level.SetLevel(logLevel)

	// Settings
	loggerConfiguration.Encoding = viper.GetString("logger.encoding")
	loggerConfiguration.Development = viper.GetBool("logger.dev_mode")
	loggerConfiguration.DisableCaller = viper.GetBool("logger.disable_caller")
	loggerConfiguration.DisableStacktrace = viper.GetBool("logger.disable_stacktrace")

	// Enable Color
	if viper.GetBool("logger.color") {
		loggerConfiguration.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Use sane timestamp when logging to console
	if loggerConfiguration.Encoding == "console" {
		loggerConfiguration.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// JSON Fields
	loggerConfiguration.EncoderConfig.MessageKey = "msg"
	loggerConfiguration.EncoderConfig.LevelKey = "level"
	loggerConfiguration.EncoderConfig.CallerKey = "caller"

	// Build the logger
	globalLogger, err := loggerConfiguration.Build()

	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(globalLogger)
	logger = globalLogger.Sugar().With("package", "cmd")

}

func initProfiler() {
	if viper.GetBool("profiler.enabled") {
		hostPort := net.JoinHostPort(viper.GetString("profiler.host"), viper.GetString("profiler.port"))
		go http.ListenAndServe(hostPort, nil)
		logger.Infof("Profiler enabled on http://%s", hostPort)
	}
}

func init() {
	cobra.OnInitialize(initConfiguration, initLogger, initProfiler)
	rootCmd.PersistentFlags().StringVarP(&configuration, "config", "c", "", "Configuration file")
}

// Execute starts the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
