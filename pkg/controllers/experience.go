package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/auth"
	"ridnvil-dev/pkg/models"
)

type ExperienceController struct {
	DB *gorm.DB
}

func NewExperienceController(db *gorm.DB) *ExperienceController {
	return &ExperienceController{DB: db}
}

func (h *ExperienceController) GetExperience(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	var experiences []models.Experiences

	if err := h.DB.Where("profile_id = ?", claim.ID).Find(&experiences).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(experiences)
}

func (h *ExperienceController) CreateExperience(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	if claim.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var experience models.Experiences

	if err := ctx.BodyParser(&experience); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	experience.ProfileID = claim.ID

	if err := h.DB.Create(&experience).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Sucess"})
}
