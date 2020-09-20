package database

import (
	"fmt"

	"github.com/PrunedNeuron/Fluoride/config"
	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// DBClient is the database client type
type DBClient struct {
	logger *zap.SugaredLogger
	conf   config.DatabaseConfiguration
	dbName string
	db     *sqlx.DB
}

// New returns a new database client
func New() (*DBClient, error) {
	logger := zap.S().With("package", "storage.postgres")
	conf := config.GetConfig().Database

	var (
		err           error
		dbCredentials string
		dbAddress     string
		dbURLOptions  string
	)

	// Username
	if username := conf.Username; username != "" {
		dbCredentials = username + ":" + conf.Password
	} else {
		return nil, fmt.Errorf("No username specified")
	}

	// Hostname + Port
	if hostname := conf.Host; hostname != "" {
		dbAddress += "@" + hostname
	} else {
		return nil, fmt.Errorf("No hostname specified")
	}

	if port := conf.Port; port != "" {
		dbAddress += ":" + port
	}

	// Database name
	dbName := conf.Database
	if dbName == "" {
		return nil, fmt.Errorf("No database specified")
	}

	// SSL mode
	if sslMode := conf.SSL; sslMode != "" {
		dbURLOptions += fmt.Sprintf("?sslmode=%s", sslMode)
	}

	// Concatenate and form the full database url
	dbURL := "postgres://" + dbCredentials + dbAddress + "/" + dbName + dbURLOptions

	// If the stop flag is caught while sleeping
	if config.Stop.Bool() {
		return nil, fmt.Errorf("Database connection aborted")
	}

	// Open a connection to the database and verify with a ping
	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %s", err)
	}

	db.SetMaxOpenConns(conf.MaxConnections)

	logger.Debugw("Connected to database server",
		"storage.host", conf.Host,
		"storage.username", conf.Username,
		"storage.port", conf.Port,
		"storage.database", conf.Database,
	)

	return &DBClient{
		logger: logger,
		conf:   conf,
		dbName: dbName,
		db:     db,
	}, nil

}
