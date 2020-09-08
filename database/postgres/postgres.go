package postgres

import (
	"fmt"
	"icon-requests/config"
	"os"

	_ "github.com/jackc/pgx/stdlib" // For pg driver
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// DBClient is the database client
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
		dbURL         string
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
		dbURL += "@" + hostname
	} else {
		return nil, fmt.Errorf("No hostname specified")
	}

	if port := viper.GetString("storage.port"); port != "" {
		dbURL += ":" + port
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

	var fullDbURL string

	// Build full DB URL
	if os.Getenv("DB_URL") != "" {
		fullDbURL = os.Getenv("DB_URL")
	} else {
		fullDbURL = "postgres://" + dbCredentials + dbURL + "/" + dbName + dbURLOptions
	}

	// If the stop flag is caught while sleeping
	if config.Stop.Bool() {
		return nil, fmt.Errorf("Database connection aborted")
	}

	// Scream if still disconnected
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %s", err)
	}

	// Connect to the database
	db, err := sqlx.Connect("pgx", fullDbURL)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %s", err)
	}

	// Ping database
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Could not ping database %s", err)
	}

	db.SetMaxOpenConns(viper.GetInt("storage.max_connections"))

	logger.Debugw("Connected to database server",
		"storage.host", viper.GetString("storage.host"),
		"storage.username", viper.GetString("storage.username"),
		"storage.port", viper.GetInt("storage.port"),
		"storage.database", viper.GetString("storage.database"),
	)

	c := &DBClient{
		logger: logger,
		dbName: dbName,
		db:     db,
		newID: func() string {
			return xid.New().String()
		},
	}

	return c, nil

	/* logger.Debug("Trying to connect...")
	db, err := sqlx.Connect("pgx", os.Getenv("DB_URL"))
	if err != nil {
		logger.Error("Could not connect to database: %s", err)
		return nil
	}

	return &DBClient{
		logger: logger,
		db:     db,
		newID: func() string {
			return xid.New().String()
		},
	} */
}

/* func (dbc *DBClient) GetDB() *sqlx.DB {
	return dbc.db
}
*/
