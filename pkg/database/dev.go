package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"

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
	err := row.StructScan(user)

	if user == (model.User{}) {
		zap.S().Errorf("Developer does not exist in the users relation")
		return false, fmt.Errorf("Developer not found")
	}

	if err != nil {
		zap.S().Errorf("Failed to retrieve developer from users table")
		return false, err
	}

	zap.S().Debugw("Querying database to check whether dev tables exist")
	_, err = dbc.db.Queryx("SELECT '$1'::regclass", fmt.Sprintf("%s_icon_packs", dev))

	if err != nil {
		zap.S().Debugw("Developer icon packs table does not exist")
		return false, err
	}

	_, err = dbc.db.Queryx("SELECT '$1'::regclass", fmt.Sprintf("%s_icon_requests", dev))

	if err != nil {
		zap.S().Debugw("Developer icon requests table does not exist")
		return false, err
	}

	return true, nil
}

// GetAllDevs gets all the icon packs
func (dbc *DBClient) GetAllDevs() ([]model.User, error) {
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

// GetPacksByDev gets all the icon packs
func (dbc *DBClient) GetPacksByDev(dev string) ([]model.Pack, error) {
	packs := []model.Pack{}
	zap.S().Debugw("Querying the database for all icon packs by given developer")

	query := fmt.Sprintf(`
		SELECT * FROM icon_packs.%s_icon_packs
		ORDER BY id DESC
	`, dev)

	rows, err := dbc.db.Queryx(query)

	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var pack model.Pack
		err = rows.StructScan(&pack)
		packs = append(packs, pack)
	}

	zap.S().Debugw("Returning with the list of all icon packs by the developer")
	return packs, nil
}

// GetIconsByDev gets all the icon packs
func (dbc *DBClient) GetIconsByDev(dev string) ([]model.Icon, error) {
	// unimplemented
	return nil, nil
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
