package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// ApplicationConfiguration is the application config model
type ApplicationConfiguration struct {
	Name    string `json:"name" yaml:"name" mapstructure:"name"`
	Author  string `json:"author" yaml:"author" mapstructure:"author"`
	Version string `json:"version" yaml:"version" mapstructure:"version"`
}

func (config *ApplicationConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *ApplicationConfiguration) IsEmpty() bool {
	return config.Version == ""
}
