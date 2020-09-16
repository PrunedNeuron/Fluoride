package database

import (
	"database/sql"
	"fluoride/internal/model"
	"fluoride/pkg/errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // For the pg driver
	"go.uber.org/zap"
)

// GetIcons gets all the icons in the DB
// !!!UNIMPLEMENTED
func (dbc *DBClient) GetIcons() ([]model.Icon, error) {
	// !!!UNIMPLEMENTED
	return nil, nil
}

// GetIconsByDev gets all the icon packs by the dev
func (dbc *DBClient) GetIconsByDev(dev string) ([]model.Icon, error) {
	icons := []model.Icon{}
	zap.S().Debugw("Querying the database for all icon requests by given developer")

	query := fmt.Sprintf(`
		SELECT * FROM icon_requests.%s_icon_requests
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
		var icon model.Icon
		err = rows.StructScan(&icon)
		icons = append(icons, icon)
		if err != nil {
			zap.S().Errorf("Failed to scan result")
			return nil, err
		}
	}

	zap.S().Debugw("Returning with the list of all icon requests belonging to the developer")
	return icons, nil
}

// GetIconsByPackByDev gets all the icons by the developer
func (dbc *DBClient) GetIconsByPackByDev(dev, pack string) ([]model.Icon, error) {

	if devExists, _ := dbc.DevExists(dev); !devExists {
		zap.S().Errorf("Developer does not exist, cannot retrieve icon requests")
		return nil, fmt.Errorf("Developer does not exist, cannot retrieve icon requests")
	}

	if packExists, _ := dbc.PackExists(dev, pack); !packExists {
		zap.S().Errorf("Icon pack does not exist, cannot retrieve icon requests")
		return nil, fmt.Errorf("Icon pack does not exist, cannot retrieve icon requests")
	}

	icons := []model.Icon{}
	zap.S().Debugw("Querying the database for all icon requests belonging to the developer")
	query := fmt.Sprintf(`
		SELECT * FROM icon_requests.%s_icon_requests
		WHERE icon_pack_name = $1
		ORDER BY id DESC
	`, dev)
	rows, err := dbc.db.Queryx(query, pack)
	zap.S().Debugw("Scanning the result")
	for rows.Next() {
		var icon model.Icon
		err = rows.StructScan(&icon)
		icons = append(icons, icon)
	}
	if err == sql.ErrNoRows {
		zap.S().Errorf("No rows in the database!")
		return nil, err
	} else if err != nil {
		zap.S().Errorf(errors.ErrDatabase.Error())
		return nil, err
	}

	zap.S().Debugw("Returning with the list of icon requests for the given icon pack")
	return icons, nil
}

// GetPendingIconsByPackByDev retrieves the list of icons which are still pending
func (dbc *DBClient) GetPendingIconsByPackByDev(dev, pack string) ([]model.Icon, error) {
	if devExists, _ := dbc.DevExists(dev); !devExists {
		zap.S().Errorf("Developer does not exist, cannot retrieve icon requests")
		return nil, fmt.Errorf("Developer does not exist, cannot retrieve icon requests")
	}

	icons := []model.Icon{}
	zap.S().Debugw("Querying the database for icons with status pending")
	query := fmt.Sprintf(`
		SELECT * FROM icon_requests.%s_icon_requests
		WHERE pack = $1 AND status = 'pending'
		ORDER BY id DESC
	`, dev)

	rows, err := dbc.db.Queryx(query, pack)
	zap.S().Debugw("Scanning the result")

	for rows.Next() {
		var icon model.Icon
		err = rows.StructScan(&icon)
		icons = append(icons, icon)
	}
	if err == sql.ErrNoRows {
		zap.S().Info("No rows in the database!")
	} else if err != nil {
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all icon requests")
	return icons, nil
}

// GetDoneIconsByPackByDev retrieves the list of icons which are still pending
func (dbc *DBClient) GetDoneIconsByPackByDev(dev, pack string) ([]model.Icon, error) {
	if devExists, _ := dbc.DevExists(dev); !devExists {
		zap.S().Errorf("Developer does not exist, cannot retrieve icon requests")
		return nil, fmt.Errorf("Developer does not exist, cannot retrieve icon requests")
	}

	icons := []model.Icon{}
	zap.S().Debugw("Querying the database for icons with status pending")

	query := fmt.Sprintf(`
		SELECT * FROM icon_requests.%s_icon_requests
		WHERE pack = $1 AND status = 'done'
		ORDER BY id DESC
	`, dev)

	rows, err := dbc.db.Queryx(query, pack)
	zap.S().Debugw("Scanning the result")

	for rows.Next() {
		var icon model.Icon
		err = rows.StructScan(&icon)
		icons = append(icons, icon)
	}
	if err == sql.ErrNoRows {
		zap.S().Info("No rows in the database!")
	} else if err != nil {
		return nil, err
	}

	zap.S().Debugw("Returning with the list of all icon requests")
	return icons, nil
}

// GetIconByComponentByPackByDev returns the matching icon
func (dbc *DBClient) GetIconByComponentByPackByDev(dev, pack, component string) (*model.Icon, error) {
	zap.S().Debugw("Querying the database with the given component")

	query := fmt.Sprintf(`
		SELECT * FROM icon_requests.%s_icon_requests
		WHERE pack = $1 AND component = $2
	`, dev)

	row := dbc.db.QueryRowx(query, pack, component)

	zap.S().Debugw("Scanning the selected icon request")
	var icon model.Icon
	err := row.StructScan(&icon)

	if err == sql.ErrNoRows {
		return nil, errors.ErrDatabaseNotFound
	} else if err != nil {
		return nil, err
	}

	zap.S().Debugw("Returning with the selected icon request")
	return &icon, nil
}

// GetIconCountByDev returns the number of icon requests owned by the dev
func (dbc *DBClient) GetIconCountByDev(dev string) (int, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*) AS COUNT FROM icon_requests.%s_icon_requests
	`, dev)
	row := dbc.db.QueryRowx(query)

	var count int
	err := row.Scan(&count)

	if err != nil {
		zap.S().Debugw("Failed to scan count")
		return -1, err
	}

	return count, nil
}

// GetPendingIconCountByDev returns the number of icon request in the database
func (dbc *DBClient) GetPendingIconCountByDev(dev string) (int, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*) AS COUNT FROM icon_requests.%s_icon_requests
		WHERE status = 'pending'
	`, dev)
	row := dbc.db.QueryRowx(query, dev)

	var count int
	err := row.Scan(&count)

	if err != nil {
		zap.S().Debugw("Failed to scan count")
		return -1, err
	}

	return count, nil
}

// GetDoneIconCountByDev returns the number of icon request in the database
func (dbc *DBClient) GetDoneIconCountByDev(dev string) (int, error) {

	query := fmt.Sprintf(`
		SELECT COUNT(*) AS COUNT FROM icon_requests.%s_icon_requests
		WHERE AND status = 'done'
	`, dev)

	row := dbc.db.QueryRowx(query)

	var count int
	err := row.Scan(&count)

	if err != nil {
		zap.S().Debugw("Failed to scan count")
		return -1, err
	}

	return count, nil
}

// SaveIcon upserts the icon to the database and updates requester count on conflict
// !UNUSED
func (dbc *DBClient) SaveIcon(dev string, icon *model.Icon) (int, error) {
	zap.S().Debugw("Upserting the given icon request into the database")

	query := fmt.Sprintf(`
		INSERT INTO icon_requests.%s_icon_requests (name, component, url, icon_pack_name)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (icon_pack_name, component) DO UPDATE
		SET requesters = icon_requests.requesters + 1
		RETURNING *
	`, dev)

	row := dbc.db.QueryRowx(query, icon.Name, icon.Component, icon.URL, icon.Pack)

	zap.S().Debugw("Scanning the inserted icon request")

	var returned model.Icon
	err := row.StructScan(&returned)

	if err != nil {
		return returned.ID, err
	}

	zap.S().Debugw("Returning with the inserted icon request ID")
	return returned.ID, nil
}

// SaveIcons upserts the list of icons to the database and updates requester counts on conflict
func (dbc *DBClient) SaveIcons(dev string, icons []*model.Icon) (int, error) {
	zap.S().Debugw("Inserting the list of icons into the database")

	for _, icon := range icons {
		zap.S().Debugw("Executing the query...")

		query := fmt.Sprintf(`
		INSERT INTO icon_requests.%s_icon_requests (name, component, url, icon_pack_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (icon_pack_name, component) DO UPDATE
		SET (requesters, updated_at) = (icon_requests.requesters + 1, CURRENT_TIMESTAMP)
	`, dev)

		_, err := dbc.db.Exec(query, icon.Name, icon.Component, icon.URL, icon.Pack, time.Now(), time.Now())

		if err != nil {
			zap.S().Debugw("Failed to insert icon")
			return -1, err
		}

	}

	zap.S().Debugw("Returning with the number of icons inserted")

	// Needs fix, updated icons also returned as inserted icon count
	return len(icons), nil
}

// UpdateIconStatus updates the status of the icon request (pending | complete)
func (dbc *DBClient) UpdateIconStatus(dev, pack, component, status string) (string, error) {
	query := fmt.Sprintf(`
		UPDATE icon_requests.%s_icon_requests
		SET status = $1
		WHERE pack = $2 AND component = $3
		RETURNING status	
	`, dev)
	row := dbc.db.QueryRowx(query, status, pack, component)

	var newStatus string
	err := row.Scan(&newStatus)

	if err != nil {
		zap.S().Debugw("Failed to scan status")
		return "", err
	}

	zap.S().Debugw("Returning with the updated status")

	return newStatus, nil
}
