package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
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

// ConfigureLogger initializes and returns the configured logger
func ConfigureLogger() (*zap.Logger, error) {
	loggerConfiguration := zap.NewProductionConfig()

	var logLevel zapcore.Level
	conf := GetConfig().Logger
	if err := logLevel.Set(conf.Level); err != nil {
		logger.Errorw("Could not determine logger.level", "error", err)
	}
	loggerConfiguration.Level.SetLevel(logLevel)

	loggerConfiguration.Encoding = conf.Encoding
	loggerConfiguration.Development = conf.DevMode
	loggerConfiguration.DisableCaller = conf.DisableCaller
	loggerConfiguration.DisableStacktrace = conf.DisableStacktrace

	if conf.Color {
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

	return loggerConfiguration.Build()
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
