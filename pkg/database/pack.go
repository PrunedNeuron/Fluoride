package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/PrunedNeuron/Fluoride/internal/model"
	"github.com/PrunedNeuron/Fluoride/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// PackExists checks whether the icon pack exists
func (dbc *DBClient) PackExists(dev, pack string) (bool, error) {
	zap.S().Debugw("Checking whether the developer exists")
	if devExists, _ := dbc.DevExists(dev); !devExists {
		zap.S().Errorf("Developer does not exist, cannot retrieve icon packs")
		return false, fmt.Errorf("Developer does not exist, cannot retrieve icon packs")
	}

	query := fmt.Sprintf(`
		SELECT * FROM icon_packs.%s_icon_packs
		WHERE name = $1
	`, dev)

	fmt.Printf("pack = %s", pack)
	_, err := dbc.db.Queryx(query, pack)

	if err == sql.ErrNoRows {
		zap.S().Debugw("Query returned no rows")
		return false, err
	}

	if err != nil {
		zap.S().Errorf("Failed to scan returned icon pack, error: %s", err.Error())
		return false, err
	}

	return true, nil
}

// CreatePack creates a new icon pack record
func (dbc *DBClient) CreatePack(pack model.Pack) (string, error) {
	zap.S().Debugw("Make sure developer exists before attempting to create a new icon pack")
	if devExists, _ := dbc.DevExists(pack.DevUsername); !devExists {
		zap.S().Errorf("Developer does not exist, cannot create icon pack")
		return "", fmt.Errorf("Developer does not exist, cannot create icon pack for non existent developer")
	}
	query := fmt.Sprintf(`
			INSERT INTO icon_packs.%s_icon_packs (name, developer_username, url, billing_status)
			VALUES ($1, $2, $3, $4)
			RETURNING name
		`, pack.DevUsername)

	zap.S().Debugw("Inserting icon pack into developer's icon packs table")
	row := dbc.db.QueryRowx(query, strings.ToLower(pack.Name), pack.DevUsername, pack.URL, pack.BillingStatus)

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

// GetPackCountByDev gets the number of icon packs by the dev
func (dbc *DBClient) GetPackCountByDev(dev string) (int, error) {

	zap.S().Debugw("Querying the database for all icon packs by given developer")

	query := fmt.Sprintf(`
		SELECT COUNT(*) FROM icon_packs.%s_icon_packs
	`, dev)

	row := dbc.db.QueryRowx(query)

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

	zap.S().Debugw("Returning with the number of icon packs by the developer")
	return count, nil
}

// GetPacks gets all the icon packs in the database
func (dbc *DBClient) GetPacks() ([]model.Pack, error) {
	devs, err := dbc.GetDevs()

	if err != nil {
		zap.S().Errorf("Failed to retrieve list of developers")
		return nil, err
	}

	var packs []model.Pack

	for _, dev := range devs {
		packsByDev, err := dbc.GetPacksByDev(dev.Username)
		if err != nil {
			zap.S().Errorf("Failed to get icon packs by developer")
			return nil, err
		}
		packs = append(packs, packsByDev...)
	}

	return packs, nil
}
