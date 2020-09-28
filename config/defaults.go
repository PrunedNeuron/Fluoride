package config

import (
	"net/http"

	"github.com/spf13/viper"
)

// SetDefaults the defaults
func SetDefaults() {

	// Application info
	viper.SetDefault("application.name", "Fluoride")
	viper.SetDefault("application.version", "1.0.0")

	// Logger Defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", true)
	viper.SetDefault("logger.disable_stacktrace", true)

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "127.0.0.1")
	viper.SetDefault("profiler.port", "3001")

	// Server Configuration
	viper.SetDefault("server.network", "tcp")
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)
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
	viper.SetDefault("storage.username", "postgres")
	viper.SetDefault("storage.password", "postgres")
	// viper.SetDefault("storage.host", "postgres")
	viper.SetDefault("storage.port", 5432)
	viper.SetDefault("storage.database", "postgres")
	viper.SetDefault("storage.ssl", "disable")
	viper.SetDefault("storage.retries", 5)
	viper.SetDefault("storage.sleep_between_retries", "5s")
	viper.SetDefault("storage.max_connections", 80)

	// Documentation settings
	viper.SetDefault("documentation.title", "API documentation")
	viper.SetDefault("documentation.base_path", "")
	viper.SetDefault("documentation.path", "/docs")
	viper.SetDefault("documentation.spec_url", "/swagger.json")
	viper.SetDefault("documentation.redoc_url", "https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js")
}
