package config

// ProfilerConfiguration is the profiler config model
type ProfilerConfiguration struct {
	Enabled bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Host    string `json:"host" yaml:"host" mapstructure:"host"`
	Port    string `json:"port" yaml:"port" mapstructure:"port"`
}
