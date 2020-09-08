package config

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// ConfigureDefaults configures the defaults
func init() {
	godotenv.Load()

	// Logger Defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", true)

	// Pidfile
	viper.SetDefault("pidfile", "")

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "")
	viper.SetDefault("profiler.port", "6060")

	// DB Defaults
	viper.SetDefault("DB_NETWORK", os.Getenv("DB_NETWORK"))
	viper.SetDefault("DB_URL", os.Getenv("DB_URL"))
	viper.SetDefault("DB_ADDRESS", os.Getenv("DB_ADDRESS"))
	viper.SetDefault("DB_USER", os.Getenv("DB_USER"))
	viper.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
	viper.SetDefault("DB_NAME", os.Getenv("DB_NAME"))

	// Server Configuration
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.tls", false)
	viper.SetDefault("server.devcert", false)
	viper.SetDefault("server.certfile", "server.crt")
	viper.SetDefault("server.keyfile", "server.key")
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
	viper.SetDefault("storage.database", "amphetamine")
	viper.SetDefault("storage.sslmode", "disable")
	viper.SetDefault("storage.retries", 5)
	viper.SetDefault("storage.sleep_between_retries", "7s")
	viper.SetDefault("storage.max_connections", 80)
	viper.SetDefault("storage.wipe_confirm", false)
}
