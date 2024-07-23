package message

import "time"

const (
	StatusSent Status = "SENT"
	StatusRead Status = "READ"
)

type (
	Status string

	Message struct {
		ID          int64     `json:"id"`
		Text        string    `json:"text"`
		UserID      int64     `json:"user_id"`
		ChatID      int64     `json:"chat_id"`
		Status      Status    `json:"status"`
		DateCreated time.Time `json:"date_created"`
		LastUpdated time.Time `json:"last_updated"`
	}
)
