package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(db *gorm.DB) *ProfileController {
	return &ProfileController{DB: db}
}

func (h *ProfileController) GetProfiles(ctx *fiber.Ctx) error {
	email := ctx.Locals("email").(string)
	var profiles models.Profile
	if errget := h.DB.First(&profiles, "email = ?", email).Error; errget != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   errget,
		})
	}
	return ctx.JSON(profiles)
}
