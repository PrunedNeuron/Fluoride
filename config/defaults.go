package config

import (
	"net/http"

	"github.com/spf13/viper"
)

// ConfigureDefaults configures the defaults
func init() {

	// Application info
	viper.SetDefault("application.version", "1.0.0")

	// Logger Defaults
	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", false)

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "localhost")
	viper.SetDefault("profiler.port", "3001")

	// Server Configuration
	viper.SetDefault("server.network", "tcp")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.log_requests", true)
	viper.SetDefault("server.log_requests_body", false)
	viper.SetDefault("server.log_disabled_http", []string{"/version"})
	viper.SetDefault("server.profiler_enabled", false)
	viper.SetDefault("server.profiler_path", "/debug")
	viper.SetDefault("server.cors.allowed_origins", []string{"*"})
	viper.SetDefault("server.cors.allowed_methods", []string{http.MethodHead, http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch})
	viper.SetDefault("server.cors.allowed_headers", []string{"*"})
	viper.SetDefault("server.cors.allowed_credentials", false)
	viper.SetDefault("server.cors.max_age", 300)

	// Database Settings
	viper.SetDefault("storage.type", "postgres")
	viper.SetDefault("storage.username", "ayush")
	viper.SetDefault("storage.password", "")
	viper.SetDefault("storage.host", "localhost")
	viper.SetDefault("storage.port", 5432)
	viper.SetDefault("storage.database", "fluoride_dev")
	viper.SetDefault("storage.sslmode", "disable")
	viper.SetDefault("storage.retries", 5)
	viper.SetDefault("storage.sleep_between_retries", "5s")
	viper.SetDefault("storage.max_connections", 80)
}
