package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ridnvil-dev/pkg/auth"
)

func GetProfiles(ctx *fiber.Ctx) error {
	user, err := auth.ProtectedHandle(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(user)
}
