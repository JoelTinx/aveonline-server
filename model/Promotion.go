package model

type Promotion struct {
	Id          int32   `json:"id" db:"id" gorm:"primaryKey"`
	Description string  `json:"description" db:"description"`
	Percentage  float64 `json:"percentage" db:"percentage"`
	StartDate   string  `json:"start_date" db:"start_date"`
	EndDate     string  `json:"end_date" db:"end_date"`
}
