package controller

import (
	"github.com/JoelTinx/aveonline-server/model"
	"gorm.io/gorm"
)

// HealthCheck godoc
// @Summary Show the list of promotions.
// @Description get the list of promotions.
// @Tags promotions
// @Accept */*
// @Produce json
// @Success 200 {array} model.Promotion
// @Router /api/promotions [get]
func GetPromotions(db *gorm.DB) ([]model.Promotion, error) {
	promotions := make([]model.Promotion, 0)
	db.Table("promotion").Find(&promotions)
	return promotions, nil
}

// HealthCheck godoc
// @Summary create a new promotion.
// @Description create a new promotion.
// @Tags promotions
// @Accept */*
// @Produce json
// @Success 200 {object} model.Promotion
// @Router /api/promotions [post]
func PostPromotion(db *gorm.DB, promotion model.Promotion) error {
	db.Table("promotion").Create(&promotion)
	return nil
}
