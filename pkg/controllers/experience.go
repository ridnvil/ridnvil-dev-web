package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type ExperienceController struct {
	DB *gorm.DB
}

func NewExperienceController(db *gorm.DB) *ExperienceController {
	return &ExperienceController{DB: db}
}

func (h *ExperienceController) GetExperience(ctx *fiber.Ctx) error {
	id := ctx.Locals("ID").(int)

	var experiences []models.Experiences
	if err := h.DB.Where("profile_id = ?", id).Find(&experiences).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(experiences)
}

func (h *ExperienceController) CreateExperience(ctx *fiber.Ctx) error {
	id := ctx.Locals("ID").(int)
	var experience models.Experiences

	if err := ctx.BodyParser(&experience); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	experience.ProfileID = id

	if err := h.DB.Create(&experience).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Sucess"})
}
