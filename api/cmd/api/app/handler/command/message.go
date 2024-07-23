package command

type (
	// MessageCommand is a command for creating a new message.
	MessageCommand struct {
		Text   string `json:"text" binding:"required"`
		UserID int64  `json:"user_id" binding:"required"`
		ChatID int64  `json:"chat_id" binding:"required"`
	}
)
