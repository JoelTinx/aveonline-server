package cmd

import (
	"log"

	"github.com/JoelTinx/aveonline-server/controller"
	"github.com/JoelTinx/aveonline-server/model"
	"github.com/JoelTinx/aveonline-server/storage"
	"github.com/gofiber/fiber/v2"
)

func GetInvoices(c *fiber.Ctx) error {
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	invoices, err := controller.GetInvoices(db)
	if err != nil {
		return err
	}
	return c.JSON(invoices)
}

func PostInvoice(c *fiber.Ctx) error {
	var invoice model.Invoice
	err := c.BodyParser(&invoice)
	if err != nil {
		return err
	}

	log.Println(invoice)
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	err = controller.PostInvoice(db, invoice)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Invoice created successfully",
	})
}
