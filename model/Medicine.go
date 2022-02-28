package model

type Medicine struct {
	Id       int32   `json:"id" db:"id" gorm:"primaryKey"`
	Name     string  `json:"name" db:"name"`
	Price    float64 `json:"price" db:"price"`
	Location string  `json:"location" db:"location"`
	// Path        string `json:"path" db:"path"`
	// Parent      string `json:"parent" db:"parent"`
	// Active      bool   `json:"active" db:"active"`
	// CreatedAt string `json:"created_at" db:"created_at"`
	// UpdatedAt string `json:"updated_at" db:"updated_at"`
	// CreatedBy string `json:"created_by" db:"created_by"`
	// UpdatedBy string `json:"updated_by" db:"updated_by"`
}
