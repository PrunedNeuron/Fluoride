package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetAllDevs gets all the icon packs
func (dbc *DBClient) GetAllDevs() ([]model.User, error) {
	devs := []model.User{}
	zap.S().Debugw("Querying the database for all icon pack developers")
	rows, err := dbc.db.Queryx(`
		SELECT * FROM users
		WHERE role = 'developer'
		ORDER BY id DESC
	`)
	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var dev model.User
		err = rows.StructScan(&dev)
		devs = append(devs, dev)
	}
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all icon pack developers")
	return devs, nil
}
