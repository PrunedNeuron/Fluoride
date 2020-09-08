package postgres

import (
	"context"
	"database/sql"
	"icon-requests/api"
	"icon-requests/database"
)

// GetIconByComponent returns the matching icon
func (dbc *DBClient) GetIconByComponent(ctx context.Context, component string) (*api.Icon, error) {
	temp := new(api.Icon)
	err := dbc.db.GetContext(ctx, temp, "SELECT * FROM icon_requests WHERE component = $1", component)
	if err == sql.ErrNoRows {
		return nil, database.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return temp, nil
}

// GetIcons gets all the icons
func (dbc *DBClient) GetIcons(ctx context.Context) ([]api.Icon, error) {
	temp := []api.Icon{}
	/* err := dbc.db.SelectContext(ctx, temp, "SELECT * FROM icon_requests") */
	rows, err := dbc.db.Queryx("SELECT * FROM icon_requests")
	for rows.Next() {
		var icon api.Icon
		err = rows.StructScan(&icon)
		temp = append(temp, icon)
	}
	if err == sql.ErrNoRows {
		// table may just be empty
	} else if err != nil {
		return nil, err
	}

	return temp, nil
}
