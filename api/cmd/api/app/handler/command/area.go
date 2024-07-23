package command

type (
	// AreaCommand is a command for creating a new area.
	AreaCommand struct {
		Name string `json:"name" binding:"required"`
	}
)
