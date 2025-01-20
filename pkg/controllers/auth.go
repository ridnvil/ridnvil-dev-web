package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Login(ctx *fiber.Ctx) error {
	var logindata struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if errparse := ctx.BodyParser(&logindata); errparse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errparse.Error()})
	}

	log.Println(logindata)
	return ctx.JSON("Login Success")
}
