package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type SocialNetworkController struct {
	DB *gorm.DB
}

func NewSocialNetworkController(db *gorm.DB) *SocialNetworkController {
	return &SocialNetworkController{DB: db}
}

func (c *SocialNetworkController) GetSocialNetworks(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)
	var socialNetworks []models.SocialNetwork

	if err := c.DB.Where("profile_id = ?", id).Find(&socialNetworks).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(socialNetworks)
}

func (c *SocialNetworkController) CreateSocialNetwork(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)
	var socialNetwork models.SocialNetwork
	if err := ctx.BodyParser(&socialNetwork); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	socialNetwork.ProfileID = id
	if err := c.DB.Create(&socialNetwork).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Sucess"})
}
