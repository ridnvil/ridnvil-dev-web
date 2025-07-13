package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"ridnvil-dev/pkg/models"
	"strconv"
)

type ProjectController struct {
	DB *gorm.DB
}

func NewProjectController(db *gorm.DB) *ProjectController {
	return &ProjectController{DB: db}
}

func (pc *ProjectController) GetProjects(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	log.Println("GetProjects email:", email)
	var projects []models.Project
	if err := pc.DB.Find(&projects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(projects)
}

func (pc *ProjectController) CreateProject(c *fiber.Ctx) error {
	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := pc.DB.Create(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Project created successfully"})
}

func (pc *ProjectController) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := pc.DB.Where("id = ?", idInt).Delete(&models.Project{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}

func (pc *ProjectController) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := pc.DB.Where("id = ?", idInt).Updates(&project).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Project updated successfully"})
}
