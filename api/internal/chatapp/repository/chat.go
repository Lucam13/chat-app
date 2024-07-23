package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lean1097/chat-back/internal/chatapp/chat"
)

const (
	getChatsQuery         = "SELECT c.id, c.from_area_id, c.to_area_id, c.date_created, c.last_updated, a.name FROM chats c JOIN areas a ON c.from_area_id = a.id"
	getChatsByAreaIDQuery = "SELECT id, from_area_id, to_area_id, date_created, last_updated FROM chats WHERE from_area_id = ?"
	saveChatQuery         = "INSERT INTO chats (from_area_id, to_area_id, date_created, last_updated) VALUES (?, ?, ?, ?)"
	deleteChatQuery       = "DELETE FROM chats WHERE id = ?"
)

type (
	chatRepository struct {
		db *sql.DB
	}

	ChatRepository interface {
		Get(ctx context.Context) (map[string][]chat.Chat, error)
		GetByAreaID(ctx context.Context, areaID int64) ([]chat.Chat, error)
		Save(ctx context.Context, fromAreaID, toAreaID int64) error
		Delete(ctx context.Context, chatID int64) error
	}
)

func NewChatRepository(db *sql.DB) ChatRepository {
	return &chatRepository{
		db: db,
	}
}

func (r chatRepository) Get(ctx context.Context) (map[string][]chat.Chat, error) {
	rows, err := r.db.QueryContext(ctx, getChatsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chatsByArea := make(map[string][]chat.Chat)
	for rows.Next() {
		var c chat.Chat
		var areaName string
		if err := rows.Scan(&c.ID, &c.FromAreaID, &c.ToAreaID, &c.DateCreated, &c.LastUpdated, &areaName); err != nil {
			return nil, err
		}

		chatsByArea[areaName] = append(chatsByArea[areaName], c)
	}

	return chatsByArea, nil
}

func (r chatRepository) GetByAreaID(ctx context.Context, areaID int64) ([]chat.Chat, error) {
	rows, err := r.db.QueryContext(ctx, getChatsByAreaIDQuery, areaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []chat.Chat
	for rows.Next() {
		var c chat.Chat
		if err := rows.Scan(&c.ID, &c.FromAreaID, &c.ToAreaID, &c.DateCreated, &c.LastUpdated); err != nil {
			return nil, err
		}

		chats = append(chats, c)
	}

	return chats, nil
}

func (r chatRepository) Save(ctx context.Context, fromAreaID, toAreaID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	now := time.Now()

	_, err = tx.ExecContext(ctx, saveChatQuery, fromAreaID, toAreaID, now, now)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r chatRepository) Delete(ctx context.Context, chatID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, deleteChatQuery, chatID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
