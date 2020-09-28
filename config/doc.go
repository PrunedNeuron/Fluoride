package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// DocumentationConfiguration is the docs config model
type DocumentationConfiguration struct {
	Title    string `json:"title" yaml:"title" mapstructure:"title"`
	BasePath string `json:"base_path" yaml:"base_path" mapstructure:"base_path"`
	Path     string `json:"path" yaml:"path" mapstructure:"path"`
	SpecURL  string `json:"spec_url" yaml:"spec_url" mapstructure:"spec_url"`
	RedocURL string `json:"redoc_url" yaml:"redoc_url" mapstructure:"redoc_url"`
}

func (config *DocumentationConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *DocumentationConfiguration) IsEmpty() bool {
	return config.SpecURL == ""
}
