package config

// Configuration is the primary configuration model
type Configuration struct {
	Application ApplicationConfiguration `json:"application" yaml:"application" mapstructure:"application"`
	Profiler    ProfilerConfiguration    `json:"profiler" yaml:"profiler" mapstructure:"profiler"`
	Logger      LoggerConfiguration      `json:"logger" yaml:"logger" mapstructure:"logger"`
	Server      ServerConfiguration      `json:"server" yaml:"server" mapstructure:"server"`
	Database    DatabaseConfiguration    `json:"database" yaml:"database" mapstructure:"database"`
}

var configuration Configuration

// GetConfig returns the global configuration
func GetConfig() *Configuration {
	return &configuration
}

// SetConfig sets the global configuration
func SetConfig(config *Configuration) {
	configuration = *config
}
