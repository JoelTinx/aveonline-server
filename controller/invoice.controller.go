package controller

import (
	"github.com/JoelTinx/aveonline-server/model"
	"gorm.io/gorm"
)

// HealthCheck godoc
// @Summary Show the list of invoices.
// @Description get the list of invoices.
// @Tags invoices
// @Accept */*
// @Produce json
// @Success 200 {array} model.Invoice
// @Router /api/invoices [get]
func GetInvoices(db *gorm.DB) ([]model.Invoice, error) {
	Invoices := make([]model.Invoice, 0)
	db.Table("invoice").Find(&Invoices)
	return Invoices, nil
}

// HealthCheck godoc
// @Summary create a new invoice.
// @Description create a new invoice.
// @Tags invoices
// @Accept */*
// @Produce json
// @Success 200 {object} model.Invoice
// @Router /api/invoices [post]
func PostInvoice(db *gorm.DB, invoice model.Invoice) error {
	db.Table("invoice").Create(&invoice)
	return nil
}
