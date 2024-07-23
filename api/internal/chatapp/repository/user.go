package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lean1097/chat-back/internal/chatapp/user"
)

const (
	getUsersQuery    = "SELECT id, username, rol, area_id, date_created, last_updated FROM users"
	getUserByIDQuery = "SELECT id, username, rol, area_id, date_created, last_updated FROM users WHERE id = ?"
	saveUserQuery    = "INSERT INTO users (username, rol, area_id, date_created, last_updated) VALUES (?, ?, ?, ?, ?)"
	deleteUserQuery  = "DELETE FROM users WHERE id = ?"
)

type (
	userRepository struct {
		db *sql.DB
	}

	UserRepository interface {
		Get(ctx context.Context) ([]user.User, error)
		GetByID(ctx context.Context, id int64) (user.User, error)
		Save(ctx context.Context, username string, rol user.Rol, areaID int64) error
		Delete(ctx context.Context, id int64) error
	}
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Get retrieves all users from the database.
func (r userRepository) Get(ctx context.Context) ([]user.User, error) {
	rows, err := r.db.QueryContext(ctx, getUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Rol, &u.AreaID, &u.DateCreated, &u.LastUpdated); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r userRepository) GetByID(ctx context.Context, userID int64) (user.User, error) {
	var user user.User
	if err := r.db.QueryRowContext(ctx, getUserByIDQuery, userID).Scan(&user.ID, &user.Username, &user.Rol, &user.AreaID, &user.DateCreated, &user.LastUpdated); err != nil {
		return user, err
	}

	return user, nil
}

func (r userRepository) Save(ctx context.Context, username string, rol user.Rol, areaID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	now := time.Now()

	_, err = tx.ExecContext(ctx, saveUserQuery, username, rol, areaID, now, now)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r userRepository) Delete(ctx context.Context, userID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, deleteUserQuery, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
