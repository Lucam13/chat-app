package user

import "time"

const (
	RolAdmin Rol = "ADMIN"
	RolUser  Rol = "USER"
)

type (
	Rol string

	User struct {
		ID          int64     `json:"id"`
		Username    string    `json:"username"`
		Rol         Rol       `json:"rol"`
		AreaID      int64     `json:"area_id"`
		DateCreated time.Time `json:"date_created"`
		LastUpdated time.Time `json:"last_updated"`
	}
)

func IsValidRol(rol Rol) bool {
	switch rol {
	case RolAdmin, RolUser:
		return true
	default:
		return false
	}
}
