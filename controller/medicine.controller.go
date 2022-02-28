package controller

import (
	"github.com/JoelTinx/aveonline-server/model"
	"gorm.io/gorm"
)

// HealthCheck godoc
// @Summary Show the list of medicines.
// @Description get the list of medicines.
// @Tags medicines
// @Accept */*
// @Produce json
// @Success 200 {array} model.Medicine
// @Router /api/medicines [get]
func GetMedicines(db *gorm.DB) ([]model.Medicine, error) {
	Medicines := make([]model.Medicine, 0)
	db.Table("medicine").Find(&Medicines)
	return Medicines, nil
}

// HealthCheck godoc
// @Summary create a new medicine.
// @Description create a new medicine.
// @Tags medicines
// @Accept */*
// @Produce json
// @Success 200 {object} model.Medicine
// @Router /api/medicines [post]
func PostMedicine(db *gorm.DB, medicine model.Medicine) error {
	db.Table("medicine").Create(&medicine)
	return nil
}
