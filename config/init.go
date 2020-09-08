package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (

	// Config and global logger
	configFile string
	logger     *zap.SugaredLogger
)

// InitLog initializes the logger
func InitLog() {

	logConfig := zap.NewProductionConfig()

	// Log Level
	var logLevel zapcore.Level
	if err := logLevel.Set(viper.GetString("logger.level")); err != nil {
		zap.S().Fatalw("Could not determine logger.level", "error", err)
	}
	logConfig.Level.SetLevel(logLevel)

	// Settings
	logConfig.Encoding = viper.GetString("logger.encoding")
	logConfig.Development = viper.GetBool("logger.dev_mode")
	logConfig.DisableCaller = viper.GetBool("logger.disable_caller")
	logConfig.DisableStacktrace = viper.GetBool("logger.disable_stacktrace")

	// Enable Color
	if viper.GetBool("logger.color") {
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Use sane timestamp when logging to console
	if logConfig.Encoding == "console" {
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// JSON Fields
	logConfig.EncoderConfig.MessageKey = "msg"
	logConfig.EncoderConfig.LevelKey = "level"
	logConfig.EncoderConfig.CallerKey = "caller"

	// Build the logger
	globalLogger, err := logConfig.Build()

	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(globalLogger)
	logger = globalLogger.Sugar().With("package", "cmd")

}
