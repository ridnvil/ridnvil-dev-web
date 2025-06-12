package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/auth"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(db *gorm.DB) *ProfileController {
	return &ProfileController{DB: db}
}

func (h *ProfileController) GetProfiles(ctx *fiber.Ctx) error {
	user, err := auth.ProtectedHandle(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(user)
}
