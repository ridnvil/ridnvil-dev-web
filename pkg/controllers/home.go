package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}
