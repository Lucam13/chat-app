package sql

import (
	"database/sql"

	internalerrors "github.com/lean1097/chat-back/internal/platform/errors"
)

type SQLDatabase struct {
	db     *sql.DB
	dbName string
}

func (sd SQLDatabase) GetDatabaseName() string {
	return sd.dbName
}

func (sd SQLDatabase) GetDatabase() *sql.DB {
	return sd.db
}

func (sd SQLDatabase) Shutdown() error {
	if sd.db != nil {
		if err := sd.db.Close(); err != nil {
			return internalerrors.ErrClosingDatabaseConnection
		}
	}

	return nil
}
