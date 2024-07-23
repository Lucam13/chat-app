package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lean1097/chat-back/internal/chatapp/message"
)

const (
	findAllQuery = "SELECT id, text, user_id, chat_id, status, date_created, last_updated FROM messages WHERE chat_id = ?"
	saveQuery    = "INSERT INTO messages (text, user_id, chat_id, status, date_created, last_updated) VALUES (?, ?, ?, ?, ?, ?)"
)

type (
	messageRepository struct {
		db *sql.DB
	}

	MessageRepository interface {
		GetByChatID(ctx context.Context, chatID int64) ([]message.Message, error)
		Save(ctx context.Context, text string, userID, chatID int64) error
	}
)

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

// GetByChatID retrieves all messages from the database for a given chat ID.
func (r messageRepository) GetByChatID(ctx context.Context, chatID int64) ([]message.Message, error) {
	rows, err := r.db.QueryContext(ctx, findAllQuery, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []message.Message
	for rows.Next() {
		var msg message.Message
		if err := rows.Scan(&msg.ID, &msg.Text, &msg.UserID, &msg.ChatID,
			&msg.Status, &msg.DateCreated, &msg.LastUpdated); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// Save saves a message to the database.
func (r messageRepository) Save(ctx context.Context, text string, userID, chatID int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	now := time.Now()

	defaultStatus := message.StatusSent

	_, err = tx.ExecContext(ctx, saveQuery, text, userID, chatID, defaultStatus, now, now)
	if err != nil {
		return err
	}

	return tx.Commit()
}
