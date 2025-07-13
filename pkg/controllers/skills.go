package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type SkillsController struct {
	DB *gorm.DB
}

func NewSkillsController(db *gorm.DB) *SkillsController {
	return &SkillsController{DB: db}
}

func (c *SkillsController) GetSkills(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	var skills []models.Skill
	if err := c.DB.Where("profile_id = ?", id).Find(&skills).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(skills)
}

func (c *SkillsController) CreateSkill(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	var skill models.Skill
	if err := ctx.BodyParser(&skill); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	skill.ProfileID = id

	if err := c.DB.Create(&skill).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Success"})
}
