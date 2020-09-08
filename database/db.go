package database

import (
	"github.com/jmoiron/sqlx"
)

// Database holds the database object
type Database struct {
	Connection *sqlx.DB
}

// Init initializes the database connection
func Init(driverName, dataSourceName string) (*Database, error) {

	db, err := sqlx.Connect(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &Database{
		Connection: db,
	}, nil
}

// Connect establishes a connection with the database
func (db *Database) connect(driverName string, dataSourceName string) {
	var err error
	db.Connection, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
}

func (db *Database) dbList(dest interface{}, query string) error {
	err := db.Connection.Select(dest, query)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) dbGet(ID string, dest interface{}, query string) error {
	err := db.Connection.Select(dest, query, ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) dbCreate(data interface{}, query string) error {
	_, err := db.Connection.Query(query, data)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) dbUpdate(ID string, data interface{}, query string) error {
	_, err := db.Connection.Query(query, ID, data)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) dbDelete(ID, query string) error {
	_, err := db.Connection.Query(query, ID)
	if err != nil {
		return err
	}
	return nil
}
