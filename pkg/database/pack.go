package database

import (
	"fluoride/internal/model"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// CreatePack creates a new icon pack record
func (dbc *DBClient) CreatePack(pack model.Pack) (string, error) {

	zap.S().Debugw("Make sure developer exists before attempting to create a new icon pack")
	if devExists, _ := dbc.DevExists(pack.DevUsername); !devExists {
		zap.S().Errorf("Developer does not exist, cannot create icon pack")
		return "", fmt.Errorf("Developer does not exist, cannot create icon pack for non existent developer")
	}
	query := fmt.Sprintf(`
			INSERT INTO icon_packs.%s_icon_packs (name, dev_username, url, billing_status)
			VALUES ($1, $2, $3, $4)
			RETURNING name
		`, pack.DevUsername)

	zap.S().Debugw("Inserting icon pack into developer's icon packs table")
	row := dbc.db.QueryRowx(query, pack.Name, pack.DevUsername, pack.URL, pack.BillingStatus)

	zap.S().Debugw("Scanning returned row")
	var packName string
	err := row.Scan(&packName)

	if err != nil {
		zap.S().Errorf("Failed to scan return value, error: %s", err)
		return "", err
	}

	zap.S().Debugw("Returning with the name of the new icon pack")
	return packName, nil

}

// GetAllPacks gets all the icon packs in the database
func (dbc *DBClient) GetAllPacks() ([]model.Pack, error) {
	// unimplemented
	return nil, nil
}
