package errors

import "errors"

// ErrDatabase is a generic database error
var ErrDatabase = errors.New("Database error")

// ErrDatabaseNotFound is the error meant for when the database doesn't exist
var ErrDatabaseNotFound = errors.New("Database not found")
