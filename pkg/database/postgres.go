package database

import (
	"fluoride/config"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// DBClient is the database client type
type DBClient struct {
	logger *zap.SugaredLogger
	dbName string
	db     *sqlx.DB
	newID  func() string
}

// New returns a new database client
func New() (*DBClient, error) {
	logger := zap.S().With("package", "database.postgres")

	var (
		err           error
		dbCredentials string
		dbAddress     string
		dbURLOptions  string
	)

	// Username
	if username := viper.GetString("storage.username"); username != "" {
		dbCredentials = username + ":" + viper.GetString("storage.password")
	} else {
		return nil, fmt.Errorf("No username specified")
	}

	// Hostname + Port
	if hostname := viper.GetString("storage.host"); hostname != "" {
		dbAddress += "@" + hostname
	} else {
		return nil, fmt.Errorf("No hostname specified")
	}

	if port := viper.GetString("storage.port"); port != "" {
		dbAddress += ":" + port
	}

	// Database name
	dbName := viper.GetString("storage.database")
	if dbName == "" {
		return nil, fmt.Errorf("No database specified")
	}

	// SSL mode
	if sslMode := viper.GetString("storage.sslmode"); sslMode != "" {
		dbURLOptions += fmt.Sprintf("?sslmode=%s", sslMode)
	}

	// Concatenate and form the full database url
	dbURL := "postgres://" + dbCredentials + dbAddress + "/" + dbName + dbURLOptions

	// If the stop flag is caught while sleeping
	if config.Stop.Bool() {
		return nil, fmt.Errorf("Database connection aborted")
	}

	// Connect to the database and verify with a ping
	db, err := sqlx.Open("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %s", err)
	}

	// Ping database
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Could not ping the database %s", err)
	}

	db.SetMaxOpenConns(viper.GetInt("storage.max_connections"))

	logger.Debugw("Connected to database server",
		"storage.host", viper.GetString("storage.host"),
		"storage.username", viper.GetString("storage.username"),
		"storage.port", viper.GetInt("storage.port"),
		"storage.database", viper.GetString("storage.database"),
	)

	return &DBClient{
		logger: logger,
		dbName: dbName,
		db:     db,
		newID: func() string {
			return xid.New().String()
		},
	}, nil

}
