package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// CreateUser inserts a new user
func (dbc *DBClient) CreateUser(user *model.User) (string, string, error) {

	zap.S().Debugw("Inserting the given user into the table")
	row, err := dbc.db.Queryx(`
		INSERT INTO users (role, name, username, email, url)
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

	zap.S().Debugw("Returning with the username and role of the added user")
	return addedUser.Name, addedUser.Role, nil
}
