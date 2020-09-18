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

var configuration config.Configuration

var (
	customConfiguration string
	logger              *zap.SugaredLogger
	rootCmd             = &cobra.Command{
		Version: config.Version,
		Use:     config.Executable,
	}
)

func initConfiguration() {
	viper.SetTypeByDefaultValue(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Failed to read configuration file!: %s ERROR: %s\n", customConfiguration, err.Error())
	}

	var configuration config.Configuration
	err := viper.Unmarshal(&configuration)
	config.SetConfig(&configuration)

	if err != nil {
		logger.Errorf("Failed to unmarshal configuration file into struct, %s", err.Error())
	}

	println("HOST = %s", config.GetConfig().Database.Host)

	if customConfiguration != "" {
		viper.SetConfigFile(customConfiguration)
		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("Failed to read custom configuration file!: %s ERROR: %s\n", customConfiguration, err.Error())

		}
	}
}

func initLogger() {

	loggerConfiguration := zap.NewProductionConfig()

	// Log Level
	var logLevel zapcore.Level
	if err := logLevel.Set(viper.GetString("logger.level")); err != nil {
		logger.Errorw("Could not determine logger.level", "error", err)
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
		loggerConfiguration.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
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
	rootCmd.PersistentFlags().StringVarP(&customConfiguration, "config", "c", "", "Configuration file (config.yml by default)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// Execute starts the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Errorf("%s", err.Error())
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
}
