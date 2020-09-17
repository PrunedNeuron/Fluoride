package database

import (
	"database/sql"
	"fmt"

	"github.com/PrunedNeuron/Fluoride/internal/model"
	"github.com/PrunedNeuron/Fluoride/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// DevExists checks whether the developer exists
func (dbc *DBClient) DevExists(dev string) (bool, error) {

	zap.S().Debugw("Querying database to check whether dev exists")

	row := dbc.db.QueryRowx(`
		SELECT * FROM secure.users
		WHERE role = 'developer' AND username = $1
	`, dev)

	var user model.User
	err := row.StructScan(&user)

	if user == (model.User{}) {
		zap.S().Errorf("Developer does not exist in the users relation")
		return false, fmt.Errorf("Developer not found")
	}

	if err != nil {
		zap.S().Errorf("Failed to retrieve developer from users table")
		return false, err
	}

	zap.S().Debugw("Querying database to check whether dev tables exist")
	query := `
		SELECT $1::regclass
	`
	_, err = dbc.db.Queryx(query, fmt.Sprintf("icon_packs.%s_icon_packs", dev))

	if err != nil {
		zap.S().Debugw("ERROR = " + err.Error())
		zap.S().Debugw("Developer icon packs table does not exist")
		return false, err
	}

	_, err = dbc.db.Queryx(query, fmt.Sprintf("icon_requests.%s_icon_requests", dev))

	if err != nil {
		zap.S().Debugw("Developer icon requests table does not exist")
		return false, err
	}

	return true, nil
}

// GetDevs gets all the icon packs
func (dbc *DBClient) GetDevs() ([]model.User, error) {
	devs := []model.User{}
	zap.S().Debugw("Querying the database for all icon pack developers")
	rows, err := dbc.db.Queryx(`
		SELECT * FROM secure.users
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

// GetDevCount gets the number of all developers
func (dbc *DBClient) GetDevCount() (int, error) {
	zap.S().Debugw("Querying the database for all icon pack developers")
	row := dbc.db.QueryRowx(`
		SELECT COUNT(*) FROM secure.users
		WHERE role = 'developer'
	`)
	zap.S().Debugw("Scanning the result")

	var count int
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return -1, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return -1, err
	}

	zap.S().Debugw("Returning with the number of all icon pack developers")
	return count, nil
}

// GetDevByUsername gets the dev with the given username
func (dbc *DBClient) GetDevByUsername(username string) (model.User, error) {

	zap.S().Debugw("Querying the database for developer with the given username")

	query := `
		SELECT * FROM secure.users
		WHERE username = $1
		ORDER BY id DESC
	`

	row := dbc.db.QueryRowx(query, username)

	zap.S().Debugw("Scanning the result")
	var dev model.User
	err := row.StructScan(&dev)

	if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return model.User{}, err
	}

	zap.S().Debugw("Returning with the matching developer")
	return dev, nil
}
