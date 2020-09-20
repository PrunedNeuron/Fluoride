package database

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// CreateSchema creates the required schema
func (dbc *DBClient) CreateSchema() error {

	// Create schema if it does not exist yet
	zap.S().Debugw("Creating schema if it does not yet exist")

	username := dbc.conf.Username

	query := fmt.Sprintf(`
			CREATE SCHEMA IF NOT EXISTS secure AUTHORIZATION %s
		`, username)

	_, err := dbc.db.Exec(query)

	if err != nil {
		zap.S().Errorf("Failed to create secure schema, error: %s", err)
		return err
	}

	query = fmt.Sprintf(`
			CREATE SCHEMA IF NOT EXISTS icon_packs AUTHORIZATION %s
		`, username)

	_, err = dbc.db.Exec(query)

	if err != nil {
		zap.S().Errorf("Failed to create icon packs schema, error: %s", err)
		return err
	}

	query = fmt.Sprintf(`
			CREATE SCHEMA IF NOT EXISTS icon_requests AUTHORIZATION %s
		`, username)

	_, err = dbc.db.Exec(query)

	if err != nil {
		zap.S().Errorf("Failed to create icon requests schema, error: %s", err)
		return err
	}

	return nil
}

// CreateTables creates the required tables
func (dbc *DBClient) CreateTables() error {
	err := dbc.CreateSchema()
	if err != nil {
		zap.S().Errorf("Failed to create schemas")
		return err
	}

	// Create tables
	// Create users table if it does not exist yet
	zap.S().Debugw("Creating users table if it does not yet exist")

	query := `
			CREATE TABLE IF NOT EXISTS secure.users (
				id SERIAL PRIMARY KEY,
				role TEXT NOT NULL,
				name TEXT NOT NULL,
				username TEXT UNIQUE NOT NULL,
				email TEXT UNIQUE NOT NULL,
				url TEXT,
				created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
			)
		`

	_, err = dbc.db.Exec(query)

	if err != nil {
		zap.S().Errorf("Failed to create users table, error: %s", err)
		return err
	}

	return nil
}
