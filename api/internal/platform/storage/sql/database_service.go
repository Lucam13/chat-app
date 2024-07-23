package sql

import (
	"database/sql"
	"fmt"
	"time"

	internalerrors "github.com/lean1097/chat-back/internal/platform/errors"
)

const (
	maxLifetimeConnection = 30 * time.Minute
	maxIdleConnections    = 10
	maxOpenConnections    = 10
)

// databaseInitializer is an interface that defines the methods for initializing a database.
type databaseInitializer interface {
	InitializeDatabase(connectionString string) (*sql.DB, error)
}

// credentials is an interface that defines the methods for getting the database credentials.
type credentials interface {
	GetUsername() string
	GetDatabaseName() string
	GetPassword() string
	GetHost() string
}

// SQLDatabaseService is a service that provides methods for managing SQL databases.
type SQLDatabaseService struct {
	credentials credentials
	initializer databaseInitializer
}

// NewSQLDatabaseService creates a new SQLDatabaseService with the given credentials and database initializer.
func NewSQLDatabaseService(c credentials, di databaseInitializer) SQLDatabaseService {
	return SQLDatabaseService{
		credentials: c,
		initializer: di,
	}
}

// StartSQLDatabase starts a new SQL database connection.
func (sds SQLDatabaseService) StartSQLDatabase() (SQLDatabase, error) {
	connectionString := sds.createConnectionString(sds.credentials)

	db, err := sds.initializer.InitializeDatabase(connectionString)
	if err != nil || db == nil {
		return SQLDatabase{}, internalerrors.ErrInitializingDatabase
	}

	sds.configureDatabase(db)

	// Ping the database to check if it is available and accesible
	if err = db.Ping(); err != nil {
		fmt.Println("Error pinging database: ", err)
		return SQLDatabase{}, internalerrors.ErrConnectingToDatabase
	}

	return SQLDatabase{
		db:     db,
		dbName: sds.credentials.GetDatabaseName(),
	}, nil
}

// configureDatabase configures the database with the given parameters.
func (sds SQLDatabaseService) configureDatabase(db *sql.DB) {
	// sets the maximum length of time that a connection can be reused
	db.SetConnMaxLifetime(maxLifetimeConnection)

	// configures the db's maximum number of idle connections
	db.SetMaxIdleConns(maxIdleConnections)

	// maximum number of open connections to the database
	db.SetMaxOpenConns(maxOpenConnections)
}

// createConnectionString creates a new connection string with the given credentials.
func (sds SQLDatabaseService) createConnectionString(c credentials) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		c.GetUsername(),
		c.GetPassword(),
		c.GetHost(),
		c.GetDatabaseName(),
	)
}
