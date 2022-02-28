package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/medicines", GetMedicines)
	api.Post("/medicines", PostMedicine)

	api.Get("/promotions", GetPromotions)
	api.Post("/promotions", PostPromotion)

	api.Get("/invoices", GetInvoices)
	api.Post("/invoices", PostInvoice)
}
