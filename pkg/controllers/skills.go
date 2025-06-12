package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/auth"
	"ridnvil-dev/pkg/models"
)

type SkillsController struct {
	DB *gorm.DB
}

func NewSkillsController(db *gorm.DB) *SkillsController {
	return &SkillsController{DB: db}
}

func (c *SkillsController) GetSkills(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	if claim.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var skills []models.Skill
	if err := c.DB.Where("profile_id = ?", claim.ID).Find(&skills).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(skills)
}

func (c *SkillsController) CreateSkill(ctx *fiber.Ctx) error {
	claim, errsec := auth.ProtectedHandle(ctx)
	if errsec != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errsec.Error()})
	}

	if claim.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var skill models.Skill
	if err := ctx.BodyParser(&skill); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	skill.ProfileID = claim.ID

	if err := c.DB.Create(&skill).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Success"})
}
