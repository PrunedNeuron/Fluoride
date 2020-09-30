package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfiguration is the logger config model
type LoggerConfiguration struct {
	Level             string `json:"level" yaml:"level" mapstructure:"level"`
	Encoding          string `json:"encoding" yaml:"encoding" mapstructure:"encoding"`
	Color             bool   `json:"color" yaml:"color" mapstructure:"color"`
	DevMode           bool   `json:"dev_mode" yaml:"dev_mode" mapstructure:"dev_mode"`
	DisableCaller     bool   `json:"disable_caller" yaml:"disable_caller" mapstructure:"disable_caller"`
	DisableStacktrace bool   `json:"disable_stacktrace" yaml:"disable_stacktrace" mapstructure:"disable_stacktrace"`
}

var conf = GetConfig().Logger

// ConfigureLogger ..
func ConfigureLogger() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})

	machineLoggerConfig := getMachineConfig()

	consoleLogger, _ := getConsoleConfig().Build()

	machineDebugging := zapcore.AddSync(logWriter(getLogPath("debug")))
	machineErrors := zapcore.AddSync(logWriter(getLogPath("error")))

	// machine encoder is for storage purposes, in production
	machineEncoder := zapcore.NewJSONEncoder(machineLoggerConfig.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewTee(
			zapcore.NewCore(machineEncoder, machineErrors, highPriority),
			zapcore.NewCore(machineEncoder, machineDebugging, lowPriority),
		),
		consoleLogger.Core(),
	)

	logger := zap.New(core)

	return logger

}

func getLogPath(suffix string) string {
	var basePath string
	if GetConfig().Application.Environment != "production" {
		basePath = "./log"
	} else {
		basePath = "/var/log"
	}
	return fmt.Sprintf(basePath+"/%s/%s_%s.log", GetConfig().Application.Name, strings.ToLower(GetConfig().Application.Name), suffix)
}

func getMachineConfig() *zap.Config {
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.DisableCaller = false
	config.DisableStacktrace = false

	return &config
}

func getConsoleConfig() *zap.Config {

	var config zap.Config

	if GetConfig().Application.Environment == "production" {
		config := getMachineConfig()
		config.DisableCaller = false
		config.DisableStacktrace = true
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return config
	}

	config = zap.NewDevelopmentConfig()

	var logLevel zapcore.Level
	if err := logLevel.Set(conf.Level); err != nil {
		logger.Errorw("Could not determine logger.level", "error", err)
	}
	config.Level.SetLevel(logLevel)

	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	if conf.Color {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// True by default
	if conf.DisableCaller {
		config.DisableCaller = conf.DisableCaller
	}

	if conf.DisableStacktrace {
		config.DisableStacktrace = conf.DisableStacktrace
	}

	return &config
}

func logWriter(logFilePath string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     30, // days
	})
}

func (config *LoggerConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *LoggerConfiguration) IsEmpty() bool {
	return config.Encoding == "" || config.Level == ""
}
