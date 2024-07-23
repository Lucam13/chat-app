package sql

import "os"

// DatabaseCredentials is a struct that represents the credentials of a database.
type DatabaseCredentials struct {
	usernameEnvKey     string
	databaseNameEnvKey string
	passwordEnvKey     string
	hostEnvKey         string
}

// BuildDatabaseCredentials builds the database credentials.
func BuildDatabaseCredentials() (dbCredentials DatabaseCredentials) {
	return DatabaseCredentials{
		usernameEnvKey:     "MYSQL_USER",
		databaseNameEnvKey: "MYSQL_DATABASE",
		passwordEnvKey:     "MYSQL_PASSWORD",
		hostEnvKey:         "MYSQL_HOST",
	}
}

// GetUsername returns the username of the database.
func (dc DatabaseCredentials) GetUsername() string {
	return os.Getenv(dc.usernameEnvKey)
}

// GetDatabaseName returns the name of the database.
func (dc DatabaseCredentials) GetDatabaseName() string {
	return os.Getenv(dc.databaseNameEnvKey)
}

// GetPassword returns the password of the database.
func (dc DatabaseCredentials) GetPassword() string {
	return os.Getenv(dc.passwordEnvKey)
}

// GetHost returns the host of the database.
func (dc DatabaseCredentials) GetHost() string {
	return os.Getenv(dc.hostEnvKey)
}
