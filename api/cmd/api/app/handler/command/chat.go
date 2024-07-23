package command

type (
	// ChatCommand is a command for creating a new chat.
	ChatCommand struct {
		FromAreaID int64 `json:"from_area_id" binding:"required"`
		ToAreaID   int64 `json:"to_area_id" binding:"required"`
	}
)
