package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// CreateUser inserts a new user
func (dbc *DBClient) CreateUser(user *model.User) (string, string, error) {

	// Create users table if it does not exist yet
	zap.S().Debugw("Creating users table if it does not yet exist")

	query := `
			CREATE TABLE IF NOT EXISTS secure.users (
				id SERIAL PRIMARY KEY,
				role TEXT NOT NULL DEFAULT "developer",
				name TEXT NOT NULL,
				username TEXT UNIQUE NOT NULL,
				email TEXT UNIQUE NOT NULL,
				url TEXT,
				created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
			)
		`

	_, err := dbc.db.Exec(query)

	if err != nil {
		zap.S().Errorf("Failed to create users table, error: %s", err)
		return "", "", err
	}

	zap.S().Debugw("Inserting the given user into the users table")
	row, err := dbc.db.Queryx(`
		INSERT INTO secure.users (role, name, username, email, url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`, user.Role, user.Name, user.Username, user.Email, user.URL)

	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return "", "", err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return "", "", err
	}

	zap.S().Debugw("Scanning the result")

	var addedUser model.User
	err = row.StructScan(&addedUser)
	if err != nil {
		zap.S().Errorf("Failed to scan return value, error: %s", err.Error())
		return "", "", err
	}

	zap.S().Debugw("Creating table for icon packs")

	// Create icon packs table and icon requests table if the user is a dev
	if user.Role == "developer" {
		// Create empty icon packs table for the given developer
		query := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS icon_packs.%s_icon_packs (
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL,
				developer_username TEXT REFERENCES secure.users(username),
				url TEXT NOT NULL,
				billing_status TEXT NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
			)
		`, user.Username)

		_, err := dbc.db.Exec(query)
		if err != nil {
			zap.S().Errorf("Failed to create new icon pack table, error: %s", err)
			return "", "", err
		}

		zap.S().Debugw("Successfully created icon packs table for the given developer")

		// Create empty icon requests table for the given developer
		query = fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS icon_requests.%s_icon_requests (
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL,
				component TEXT UNIQUE NOT NULL,
				url TEXT NOT NULL,
				requesters TEXT NOT NULL,
				status TEXT NOT NULL,
				icon_pack_id INT REFERENCES icon_packs.%s_icon_packs(id)
				created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
			)
		`, user.Username, user.Username)

		_, err = dbc.db.Exec(query)
		if err != nil {
			zap.S().Errorf("Failed to create new icon requests table, error: %s", err)
			return "", "", err
		}

		zap.S().Debugw("Successfully created icon requests table for the given developer")
	}

	zap.S().Debugw("Returning with the username and role of the added user")
	return addedUser.Name, addedUser.Role, nil
}
