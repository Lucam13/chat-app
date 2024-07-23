package chat

import "time"

type (
	Chat struct {
		ID          int64     `json:"id"`
		FromAreaID  int64     `json:"from_area_id"`
		ToAreaID    int64     `json:"to_area_id"`
		DateCreated time.Time `json:"date_created"`
		LastUpdated time.Time `json:"last_updated"`
	}
)
