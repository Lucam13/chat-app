package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DatabaseInitializer is a struct that implements the databaseInitializer interface.
type DatabaseInitializer struct {
	opener sqlDatabaseOpener
}

// sqlDatabaseOpener is a function type that defines the signature of a function that opens a SQL database connection.
type sqlDatabaseOpener func(string, string) (*sql.DB, error)

// NewDatabaseInitializer creates a new DatabaseInitializer with the given opener.
func NewDatabaseInitializer() DatabaseInitializer {
	return DatabaseInitializer{
		opener: sql.Open,
	}
}

// InitializeDatabase initializes the database.
func (di DatabaseInitializer) InitializeDatabase(connectionString string) (*sql.DB, error) {
	return di.opener("mysql", connectionString)
}
