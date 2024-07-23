package area

import "time"

const (
	AreaNameWomenNursing        = "WOMEN_NURSING"
	AreaNameMenNursing          = "MEN_NURSING"
	AreaNameAguaribayNursing    = "AGUARIBAY_NURSING"
	AreaNameHomeDayCenter       = "HOME_DAY_CENTER"
	AreaNameAmbulatoryDayCenter = "AMBULATORY_DAY_CENTER"
	AreaNameCare                = "CARE"
	AreaNameAdministration      = "ADMINISTRATION"
	AreaNameLaundryAndWardrobe  = "LAUNDRY_AND_WARDROBE"
	AreaNameMaintenance         = "MAINTENANCE"
	AreaNamePharmacy            = "PHARMACY"
	AreaNameDoctorsOffice       = "DOCTORS_OFFICE"
)

type (
	Name string

	Area struct {
		ID          int64     `json:"id"`
		Name        Name      `json:"name"`
		DateCreated time.Time `json:"date_created"`
		LastUpdated time.Time `json:"last_updated"`
	}
)

func IsValidAreaName(name Name) bool {
	switch name {
	case AreaNameWomenNursing, AreaNameMenNursing, AreaNameAguaribayNursing, AreaNameHomeDayCenter,
		AreaNameAmbulatoryDayCenter, AreaNameCare, AreaNameAdministration, AreaNameLaundryAndWardrobe,
		AreaNameMaintenance, AreaNamePharmacy, AreaNameDoctorsOffice:
		return true
	default:
		return false
	}
}
