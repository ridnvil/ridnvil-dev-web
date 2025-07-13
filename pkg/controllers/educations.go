package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type EducationsController struct {
	DB *gorm.DB
}

func NewEducationsController(db *gorm.DB) *EducationsController {
	return &EducationsController{DB: db}
}

func (h *EducationsController) GetEducations(ctx *fiber.Ctx) error {
	id := ctx.Locals("ID").(int)
	var educations []models.Educations
	if err := h.DB.Where("profile_id = ?", id).Find(&educations).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(educations)
}

func (h *EducationsController) CreateEducation(ctx *fiber.Ctx) error {
	id := ctx.Locals("ID").(int)
	var education models.Educations
	if err := ctx.BodyParser(&education); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	education.ProfileID = id

	if err := h.DB.Create(&education).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Success"})
}
