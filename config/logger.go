package config

// LoggerConfiguration is the logger config model
type LoggerConfiguration struct {
	Level             string `json:"level" yaml:"level" mapstructure:"level"`
	Encoding          string `json:"encoding" yaml:"encoding" mapstructure:"encoding"`
	Color             bool   `json:"color" yaml:"color" mapstructure:"color"`
	DevMode           bool   `json:"dev_mode" yaml:"dev_mode" mapstructure:"dev_mode"`
	DisableCaller     bool   `json:"disable_caller" yaml:"disable_caller" mapstructure:"disable_caller"`
	DisableStacktrace bool   `json:"disable_stacktrace" yaml:"disable_stacktrace" mapstructure:"disable_stacktrace"`
}
