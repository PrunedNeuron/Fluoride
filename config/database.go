package config

// DatabaseConfiguration is the database config model
type DatabaseConfiguration struct {
	Type                string `json:"type" yaml:"type" mapstructure:"type"`
	Username            string `json:"username" yaml:"username" mapstructure:"username"`
	Password            string `json:"password" yaml:"password" mapstructure:"password"`
	Host                string `json:"host" yaml:"host" mapstructure:"host"`
	Port                string `json:"port" yaml:"port" mapstructure:"port"`
	Database            string `json:"database" yaml:"database" mapstructure:"database"`
	SSL                 string `json:"ssl" yaml:"ssl" mapstructure:"ssl"`
	Retries             int    `json:"retries" yaml:"retries" mapstructure:"retries"`
	SleepBetweenRetries string `json:"sleep_between_retries" yaml:"sleep_between_retries" mapstructure:"sleep_between_retries"`
	MaxConnections      int    `json:"max_connections" yaml:"max_connections" mapstructure:"max_connections"`
}
