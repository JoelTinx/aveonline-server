package cmd

import (
	"log"

	"github.com/JoelTinx/aveonline-server/controller"
	"github.com/JoelTinx/aveonline-server/model"
	"github.com/JoelTinx/aveonline-server/storage"
	"github.com/gofiber/fiber/v2"
)

func GetMedicines(c *fiber.Ctx) error {
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	categories, err := controller.GetMedicines(db)
	if err != nil {
		return err
	}
	return c.JSON(categories)
}

func PostMedicine(c *fiber.Ctx) error {
	var medicine model.Medicine
	err := c.BodyParser(&medicine)
	if err != nil {
		return err
	}

	log.Println(medicine)
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	err = controller.PostMedicine(db, medicine)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Medicine created successfully",
	})
}
