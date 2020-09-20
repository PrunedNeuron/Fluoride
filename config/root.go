package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Configuration is the primary configuration model
type Configuration struct {
	Application ApplicationConfiguration `json:"application" yaml:"application" mapstructure:"application"`
	Profiler    ProfilerConfiguration    `json:"profiler" yaml:"profiler" mapstructure:"profiler"`
	Logger      LoggerConfiguration      `json:"logger" yaml:"logger" mapstructure:"logger"`
	Server      ServerConfiguration      `json:"server" yaml:"server" mapstructure:"server"`
	Database    DatabaseConfiguration    `json:"database" yaml:"database" mapstructure:"database"`
}

var (
	configuration Configuration
	logger        *zap.SugaredLogger
)

// Configure loads the configuration file into the config structs
func Configure(configurations ...string) {

	logger := zap.S().With("package", "config")

	SetDefaults()

	viper.SetTypeByDefaultValue(true)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if len(configurations) > 1 {
		logger.Errorf("Cannot process multiple configuration files")
	} else if len(configurations) == 0 || configurations[0] == "" {
		logger.Debugf("No configuration file passed, using defaults") // Default configuration file if nothing is passed

		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("Failed to read configuration file! ERROR: %s\n", err.Error())
		}

		logger.Debugf("Unmarshalling the read configuration file into the configuration struct")

		err := viper.Unmarshal(&configuration)
		if err != nil {
			logger.Errorf("Failed to unmarshal configuration file into struct, %s", err.Error())
		}
	} else {
		logger.Debugf("Valid configuration file passed, using")
		viper.SetConfigFile(configurations[0])
		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("Failed to read custom configuration file!: %s ERROR: %s\n", configurations[0], err.Error())

		}
	}
}

// GetConfig returns the global configuration
func GetConfig() *Configuration {
	if configuration.IsEmpty() {
		Configure()
	}
	return &configuration
}

// SetConfig sets the global configuration
func SetConfig(config *Configuration) {
	configuration = *config
}

func (config *Configuration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *Configuration) IsEmpty() bool {
	return config.Logger.IsEmpty() && config.Server.IsEmpty() && config.Application.IsEmpty() && config.Database.IsEmpty()
}
