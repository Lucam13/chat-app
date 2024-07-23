package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lean1097/chat-back/internal/chatapp/area"
)

const (
	getAreasQuery    = "SELECT id, name, date_created, last_updated FROM areas"
	getAreaByIDQuery = "SELECT id, name, date_created, last_updated FROM areas WHERE id = ?"
	saveAreaQuery    = "INSERT INTO areas (name, date_created, last_updated) VALUES (?, ?, ?)"
	deleteAreaQuery  = "DELETE FROM areas WHERE id = ?"
)

type (
	areaRepository struct {
		db *sql.DB
	}

	AreaRepository interface {
		Get(ctx context.Context) ([]area.Area, error)
		GetByID(ctx context.Context, id int64) (area.Area, error)
		Save(ctx context.Context, areaName string) error
		Delete(ctx context.Context, id int64) error
	}
)

func NewAreaRepository(db *sql.DB) AreaRepository {
	return &areaRepository{
		db: db,
	}
}

// Get retrieves all areas from the database.
func (r areaRepository) Get(ctx context.Context) ([]area.Area, error) {
	rows, err := r.db.QueryContext(ctx, getAreasQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []area.Area
	for rows.Next() {
		var a area.Area
		if err := rows.Scan(&a.ID, &a.Name, &a.DateCreated, &a.LastUpdated); err != nil {
			return nil, err
		}
		areas = append(areas, a)
	}

	return areas, nil
}

// GetByID retrieves an area from the database by its ID.
func (r areaRepository) GetByID(ctx context.Context, id int64) (area.Area, error) {
	var a area.Area
	if err := r.db.QueryRowContext(ctx, getAreaByIDQuery, id).Scan(&a.ID, &a.Name, &a.DateCreated, &a.LastUpdated); err != nil {
		return area.Area{}, err
	}

	return a, nil
}

// Save saves an area to the database.
func (r areaRepository) Save(ctx context.Context, areaName string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	now := time.Now()

	_, err = tx.ExecContext(ctx, saveAreaQuery, areaName, now, now)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Delete deletes an area from the database by its ID.
func (r areaRepository) Delete(ctx context.Context, areaID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, deleteAreaQuery, areaID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
