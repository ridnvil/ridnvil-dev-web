package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"ridnvil-dev/pkg/models"
	"ridnvil-dev/pkg/services"
)

func CreateClient(ctx *fiber.Ctx) error {
	var ipinfo models.IPInfo

	if err := ctx.BodyParser(&ipinfo); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := services.CreateClient(ipinfo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "Created!"})
}
