package model

type Invoice struct {
	Id           int32   `json:"id" db:"id" gorm:"primaryKey"`
	CreationDate string  `json:"creation_date" db:"creation_date"`
	Total        float64 `json:"total" db:"total"`
}
