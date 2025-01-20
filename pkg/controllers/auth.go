package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"ridnvil-dev/pkg/auth"
	"ridnvil-dev/pkg/models"
	"time"
)

func Login(ctx *fiber.Ctx) error {
	var logindata struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if errparse := ctx.BodyParser(&logindata); errparse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errparse.Error()})
	}

	if logindata.Email == "rid.nvil17@gmail.com" && logindata.Password == "v5jd7Sk61lL99Yumna" {

		var user models.User

		user.ID = 1
		user.Name = "Ridwan"
		user.Email = "rid.nvil17@gmail.com"

		var responseData struct {
			User  models.User `json:"user"`
			Token string      `json:"token"`
		}

		token, err := auth.GenerateJWT(user)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		ctx.Cookie(&fiber.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour),
		})

		responseData.User = user
		responseData.Token = token
		return ctx.JSON(responseData)
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errors.New("login failed")})
}
