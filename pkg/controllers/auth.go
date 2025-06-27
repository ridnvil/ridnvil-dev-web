package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/auth"
	"ridnvil-dev/pkg/models"
	"time"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func (h *AuthController) Login(ctx *fiber.Ctx) error {
	var logindata struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if errparse := ctx.BodyParser(&logindata); errparse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errparse.Error()})
	}

	// v5jd7Sk61lL99Yumna
	if logindata.Email == "rid.nvil17@gmail.com" && logindata.Password == "M1r34cl3" {

		var user models.Profile

		if errgetprof := h.DB.First(&user, "email = ?", logindata.Email).Error; errgetprof != nil {
			if errors.Is(errgetprof, gorm.ErrRecordNotFound) {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": errgetprof.Error()})
		}

		var responseData struct {
			User models.Profile `json:"user"`
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
		return ctx.JSON(responseData)
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errors.New("login failed")})
}

func (h *AuthController) Logout(ctx *fiber.Ctx) error {
	user, err := auth.ProtectedHandle(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
