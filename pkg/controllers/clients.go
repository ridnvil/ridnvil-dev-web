package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"ridnvil-dev/pkg/models"
	"ridnvil-dev/pkg/repository"
	"ridnvil-dev/pkg/services"
)

type ClientController struct {
	DB *gorm.DB
}

func NewClientController(db *gorm.DB) *ClientController {
	return &ClientController{DB: db}
}

func (h *ClientController) CreateClient(ctx *fiber.Ctx) error {
	var ipinfo models.IPInfo

	if err := ctx.BodyParser(&ipinfo); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if err := services.CreateClient(ipinfo, h.DB); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "Created!"})
}

func (h *ClientController) GetListClient(ctx *fiber.Ctx) error {
	client, err := repository.GetListClient(h.DB)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(client)
}
