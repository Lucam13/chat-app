package errors

import "errors"

var (
	// ErrClosingDatabaseConnection is the error returned when there is an error closing the database connection.
	ErrClosingDatabaseConnection = errors.New("storage: error closing database connection")
	// ErrInitializingDatabase is the error returned when there is an error initializing the database.
	ErrInitializingDatabase = errors.New("storage: error initializing database")
	// ErrConnectingToDatabase is the error returned when there is an error connecting to the database.
	ErrConnectingToDatabase = errors.New("storage: error connecting to database")

	// ErrResourceNotFound is the error returned when a resource is not found.
	ErrResourceNotFound = errors.New("storage: resource not found")
)
