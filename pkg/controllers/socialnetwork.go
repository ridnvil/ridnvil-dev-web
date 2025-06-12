package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/auth"
	"ridnvil-dev/pkg/models"
)

type SocialNetworkController struct {
	DB *gorm.DB
}

func NewSocialNetworkController(db *gorm.DB) *SocialNetworkController {
	return &SocialNetworkController{DB: db}
}

func (c *SocialNetworkController) GetSocialNetworks(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	var socialNetworks []models.SocialNetwork

	if err := c.DB.Where("profile_id = ?", claim.ID).Find(&socialNetworks).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(socialNetworks)
}

func (c *SocialNetworkController) CreateSocialNetwork(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	if claim.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var socialNetwork models.SocialNetwork
	if err := ctx.BodyParser(&socialNetwork); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	socialNetwork.ProfileID = claim.ID
	if err := c.DB.Create(&socialNetwork).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Sucess"})
}
