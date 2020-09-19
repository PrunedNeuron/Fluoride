package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// ProfilerConfiguration is the profiler config model
type ProfilerConfiguration struct {
	Enabled bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Host    string `json:"host" yaml:"host" mapstructure:"host"`
	Port    string `json:"port" yaml:"port" mapstructure:"port"`
}

func (config *ProfilerConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}
