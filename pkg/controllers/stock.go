package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ridnvil-dev/pkg/repository"
)

func GetStock(ctx *fiber.Ctx) error {
	stock, err := repository.GetListStock()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stock)
}
