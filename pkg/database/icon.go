package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"time"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetAllIcons gets all the icons
func (dbc *DBClient) GetAllIcons() ([]model.Icon, error) {
	temp := []model.Icon{}
	zap.S().Debugw("Querying the database for all icon request...")
	rows, err := dbc.db.Queryx("SELECT * FROM icon_requests ORDER BY id DESC")
	zap.S().Debugw("Scanning the result...")
	for rows.Next() {
		var icon model.Icon
		err = rows.StructScan(&icon)
		temp = append(temp, icon)
	}
	if err == sql.ErrNoRows {
		zap.S().Info("No rows in the database!")
	} else if err != nil {
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all icon requests...")
	return temp, nil
}

// GetIconByComponent returns the matching icon
func (dbc *DBClient) GetIconByComponent(component string) (*model.Icon, error) {
	zap.S().Debugw("Querying the database with the given component...")
	row := dbc.db.QueryRowx("SELECT * FROM icon_requests WHERE component = $1", component)

	zap.S().Debugw("Scanning the selected icon request...")
	var icon model.Icon
	err := row.StructScan(&icon)

	if err == sql.ErrNoRows {
		return nil, errors.ErrDatabaseNotFound
	} else if err != nil {
		return nil, err
	}

	zap.S().Debugw("Returning with the selected icon request...")
	return &icon, nil
}

// SaveIcon upserts the icon to the database and updates requester count on conflict
// !UNUSED
func (dbc *DBClient) SaveIcon(icon *model.Icon) (int, error) {
	zap.S().Debugw("Upserting the given icon request into the database...")
	row := dbc.db.QueryRowx(`
		INSERT INTO icon_requests (name, component, url)
		VALUES ($1, $2, $3)
		ON CONFLICT (component) DO UPDATE
		SET requesters = icon_requests.requesters + 1
		RETURNING *
	`, icon.Name, icon.Component, icon.URL)

	zap.S().Debugw("Scanning the inserted icon request...")

	var returned model.Icon
	err := row.StructScan(&returned)

	if err != nil {
		return returned.ID, err
	}

	zap.S().Debugw("Returning with the inserted icon request ID...")
	return returned.ID, nil
}

// SaveIcons upserts the list of icons to the database and updates requester counts on conflict
func (dbc *DBClient) SaveIcons(icons []*model.Icon) (int, error) {
	zap.S().Debugw("Inserting the list of icons into the database...")

	for _, icon := range icons {
		zap.S().Debugw("Executing the query...")
		_, err := dbc.db.Exec(`
		INSERT INTO icon_requests (name, component, url, pack, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (component) DO UPDATE
		SET (requesters, updated_at) = (icon_requests.requesters + 1, CURRENT_TIMESTAMP)
	`, icon.Name, icon.Component, icon.URL, icon.Pack, time.Now(), time.Now())

		if err != nil {
			zap.S().Debugw("Failed to insert icon")
			return -1, err
		}

	}

	zap.S().Debugw("Returning...")

	// Needs fix, updated icons also returned as inserted icon count
	return len(icons), nil
}

// GetCount returns the number of icon request in the database
func (dbc *DBClient) GetCount() (int, error) {
	row := dbc.db.QueryRowx("SELECT COUNT(*) AS COUNT FROM icon_requests")

	var count int
	err := row.Scan(&count)

	if err != nil {
		zap.S().Debugw("Failed to scan count...")
		return -1, err
	}

	return count, nil
}

// UpdateStatus updates the status of the icon request (pending | complete)
func (dbc *DBClient) UpdateStatus(component, status string) (string, error) {
	row := dbc.db.QueryRowx(`
		UPDATE icon_requests
		SET status = $1
		WHERE component = $2
		RETURNING status	
	`, status, component)

	var newStatus string
	err := row.Scan(&newStatus)

	if err != nil {
		zap.S().Debugw("Failed to scan status...")
		return "", err
	}

	return newStatus, nil
}
