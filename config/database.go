package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// DatabaseConfiguration is the database config model
type DatabaseConfiguration struct {
	SleepBetweenRetries string `json:"sleep_between_retries" yaml:"sleep_between_retries" mapstructure:"sleep_between_retries"`
	MaxConnections      int    `json:"max_connections" yaml:"max_connections" mapstructure:"max_connections"`
	Username            string `json:"username" yaml:"username" mapstructure:"username"`
	Password            string `json:"password" yaml:"password" mapstructure:"password"`
	Database            string `json:"database" yaml:"database" mapstructure:"database"`
	Retries             int    `json:"retries" yaml:"retries" mapstructure:"retries"`
	Type                string `json:"type" yaml:"type" mapstructure:"type"`
	Host                string `json:"host" yaml:"host" mapstructure:"host"`
	Port                string `json:"port" yaml:"port" mapstructure:"port"`
	SSL                 string `json:"ssl" yaml:"ssl" mapstructure:"ssl"`
}

func (config *DatabaseConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *DatabaseConfiguration) IsEmpty() bool {
	return config.Type == "" || config.Host == "" || config.Port == "" || config.Password == "" || config.Database == ""
}
