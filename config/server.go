package config

import (
	"encoding/json"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"go.uber.org/zap"
)

// CORSConfiguration is the model for the CORS config
type CORSConfiguration struct {
	AllowedOrigins     []string `json:"allowed_origins" yaml:"allowed_origins" mapstructure:"allowed_origins"`
	AllowedMethods     []string `json:"allowed_methods" yaml:"allowed_methods" mapstructure:"allowed_methods"`
	AllowedHeaders     []string `json:"allowed_headers" yaml:"allowed_headers" mapstructure:"allowed_headers"`
	AllowedCredentials bool     `json:"allowed_credentials" yaml:"allowed_credentials" mapstructure:"allowed_credentials"`
	MaxAge             int      `json:"max_age" yaml:"max_age" mapstructure:"max_age"`
}

// ServerConfiguration is the server config model
type ServerConfiguration struct {
	Network         string            `json:"network" yaml:"network" mapstructure:"network"`
	Host            string            `json:"host" yaml:"host" mapstructure:"host"`
	Port            string            `json:"port" yaml:"port" mapstructure:"port"`
	LogRequests     bool              `json:"log_requests" yaml:"log_requests" mapstructure:"log_requests"`
	LogRequestsBody bool              `json:"log_requests_body" yaml:"log_requests_body" mapstructure:"log_requests_body"`
	LogDisabledHTTP []string          `json:"log_disabled_http" yaml:"log_disabled_http" mapstructure:"log_disabled_http"`
	ProfilerEnabled bool              `json:"profiler_enabled" yaml:"profiler_enabled" mapstructure:"profiler_enabled"`
	ProfilerPath    string            `json:"profiler_path" yaml:"profiler_path" mapstructure:"profiler_path"`
	CORS            CORSConfiguration `json:"cors" yaml:"cors" mapstructure:"cors"`
}

func (config *ServerConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

func (config *CORSConfiguration) String() string {
	json, err := json.Marshal(config)
	if err != nil {
		zap.S().Errorf("Failed to marshal struct into json, error: %s\n", err)
		return errors.ErrMarshal.Error()
	}
	return fmt.Sprintf(string(json))
}

// IsEmpty checks whether the struct is empty / uninitialized / nil
func (config *ServerConfiguration) IsEmpty() bool {
	return config.Network == "" || config.Host == "" || config.Port == ""
}
