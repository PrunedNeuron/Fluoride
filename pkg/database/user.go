package database

import (
	"database/sql"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/pkg/errors"
	"github.com/PrunedNeuron/Fluoride/pkg/model"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetUsers gets the list of all users in the DB
func (dbc *DBClient) GetUsers() ([]model.User, error) {
	users := []model.User{}
	zap.S().Debugw("Querying the database for all users")
	rows, err := dbc.db.Queryx(`
		SELECT * FROM secure.users
		ORDER BY id DESC
	`)
	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var user model.User
		err = rows.StructScan(&user)
		users = append(users, user)
	}
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all users")
	return users, nil
}

// CreateUser inserts a new user
func (dbc *DBClient) CreateUser(user *model.User) (string, string, error) {

	// Create users table if it does not exist yet

	zap.S().Debugw("Querying database to check whether users table exists")
	_, err := dbc.db.Queryx("SELECT 'secure.users'::regclass")

	if err != nil {
		zap.S().Debugw("Users table does not exist, creating")
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

		_, err := dbc.db.Exec(query)

		if err != nil {
			zap.S().Errorf("Failed to create users table, error: %s", err)
			return "", "", err
		}
	}

	zap.S().Debugw("Inserting the given user into the users table")
	row := dbc.db.QueryRowx(`
		INSERT INTO secure.users (role, name, username, email, url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`, user.Role, user.Name, user.Username, user.Email, user.URL)

	zap.S().Debugw("Scanning the result")

	var addedUser model.User
	err = row.StructScan(&addedUser)

	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return "", "", err
	} else if err != nil {
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
				name TEXT UNIQUE NOT NULL,
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
				component TEXT NOT NULL,
				url TEXT NOT NULL,
				requesters INT NOT NULL DEFAULT 0,
				status TEXT NOT NULL DEFAULT 'pending',
				icon_pack_name TEXT REFERENCES icon_packs.%s_icon_packs(name),
				created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT %s_unique_component_pack_id UNIQUE (component, icon_pack_name)
			)
		`, user.Username, user.Username, user.Username)

		_, err = dbc.db.Exec(query)
		if err != nil {
			zap.S().Errorf("Failed to create new icon requests table, error: %s", err)
			return "", "", err
		}

		zap.S().Debugw("Successfully created icon requests table for the given developer")
	}

	zap.S().Debugw("Returning with the username and role of the added user")
	return addedUser.Username, addedUser.Role, nil
}
