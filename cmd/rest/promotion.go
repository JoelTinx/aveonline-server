package cmd

import (
	"log"

	"github.com/JoelTinx/aveonline-server/controller"
	"github.com/JoelTinx/aveonline-server/model"
	"github.com/JoelTinx/aveonline-server/storage"
	"github.com/gofiber/fiber/v2"
)

func GetPromotions(c *fiber.Ctx) error {
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	promotions, err := controller.GetPromotions(db)
	if err != nil {
		return err
	}
	return c.JSON(promotions)
}

func PostPromotion(c *fiber.Ctx) error {
	var promotion model.Promotion
	err := c.BodyParser(&promotion)
	if err != nil {
		return err
	}

	log.Println(promotion)
	db, err := storage.ConnectDB()
	if err != nil {
		return err
	}
	err = controller.PostPromotion(db, promotion)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Promotion created successfully",
	})
}
