package database

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetIconPackIDFromName checks whether the plans table exists
func (dbc *DBClient) GetIconPackIDFromName(dev, pack string) (int, error) {
	zap.S().Debugw("Querying database to check whether icon pack exists")

	query := fmt.Sprintf(`
		SELECT id FROM icon_packs.%s_icon_packs
		WHERE LOWER(name) = LOWER($1)
	`, dev)

	row := dbc.db.QueryRowx(query, pack)

	var id int
	err := row.Scan(&id)

	if err != nil {
		zap.S().Errorf("Failed to scan icon pack ID, error: " + err.Error())
		return -1, err
	}

	zap.S().Debugf("Scanned id = %d", id)

	return id, nil
}
