package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ridnvil-dev/pkg/models"
)

type HomeController struct {
	DB *gorm.DB
}

func NewHomeController(db *gorm.DB) *HomeController {
	return &HomeController{DB: db}
}

func (h *HomeController) Home(ctx *fiber.Ctx) error {
	var profile models.Profile
	var profileResponse models.ProfileHome
	id := 1
	if err := h.DB.First(&profile, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Profile not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	profileResponse.ID = profile.ID
	profileResponse.Name = profile.Name
	profileResponse.Position = profile.Position
	profileResponse.Email = profile.Email
	profileResponse.Phone = profile.Phone
	profileResponse.Address = profile.Address
	profileResponse.Bio = profile.Bio

	var socialLinks []models.SocialNetwork
	if err := h.DB.Where("profile_id = ?", profile.ID).Find(&socialLinks).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	profileResponse.SocialNetworkRef = socialLinks

	var skills []models.Skill
	if err := h.DB.Where("profile_id = ?", profile.ID).Find(&skills).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	profileResponse.SkillRef = skills

	var experiences []models.Experiences
	if err := h.DB.Where("profile_id = ?", profile.ID).Find(&experiences).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	profileResponse.ExperincesRef = experiences

	var educations []models.Educations
	if err := h.DB.Where("profile_id = ?", profile.ID).Find(&educations).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	profileResponse.EducationRef = educations

	return ctx.JSON(profileResponse)
}
