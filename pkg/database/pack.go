package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetAllPacks gets all the icon packs
func (dbc *DBClient) GetAllPacks(dev string) ([]model.Pack, error) {
	packs := []model.Pack{}
	zap.S().Debugw("Querying the database for all icon packs")
	rows, err := dbc.db.Queryx(`
		SELECT * FROM icon_packs
		WHERE developer_username = $1
		ORDER BY id DESC
	`, dev)
	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var pack model.Pack
		err = rows.StructScan(&pack)
		packs = append(packs, pack)
	}
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all icon packs")
	return packs, nil
}
