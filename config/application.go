package config

// ApplicationConfiguration is the application config model
type ApplicationConfiguration struct {
	Version string `json:"version" yaml:"version" mapstructure:"version"`
}
