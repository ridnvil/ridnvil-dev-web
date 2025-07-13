package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"ridnvil-dev/pkg/controllers"
	"ridnvil-dev/pkg/database"
	"ridnvil-dev/pkg/middleware"
	"ridnvil-dev/pkg/models"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "ridnvil.dev"})
	app.Use(logger.New())

	app.Static("/", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/login", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/profile", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/experiences", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/skills", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/projects", "ridnvil/build", fiber.Static{Browse: false})

	if errdbcheck := database.AutoCreateDatabase(); errdbcheck != nil {
		log.Println(errdbcheck)
	}

	db, err := database.OpenConnectionSQLite()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&models.Profile{},
		&models.IPInfo{},
		&models.Project{},
		&models.Experiences{},
		&models.SocialNetwork{},
		&models.Educations{},
		&models.Skill{},
	); err != nil {
		panic(err)
	}

	api := app.Group("/api")

	authController := controllers.NewAuthController(db)
	clientController := controllers.NewClientController(db)
	profileController := controllers.NewProfileController(db)
	homeController := controllers.NewHomeController(db)
	experienceController := controllers.NewExperienceController(db)
	projectController := controllers.NewProjectController(db)
	taskController := controllers.NewTasksController(db)

	// Testing Push
	api.Post("/login", authController.Login)
	api.Post("/logout", middleware.JWTMiddleware(), authController.Logout)

	api.Get("/welcome", homeController.Home)
	api.Post("/client", clientController.CreateClient)
	api.Get("/client", clientController.GetListClient)

	api.Get("/profile", middleware.JWTMiddleware(), profileController.GetProfiles)

	api.Get("/experiences", middleware.JWTMiddleware(), experienceController.GetExperience)
	api.Post("/experiences", middleware.JWTMiddleware(), experienceController.CreateExperience)

	api.Get("/projects/", middleware.JWTMiddleware(), projectController.GetProjects)
	api.Post("/projects", middleware.JWTMiddleware(), projectController.CreateProject)
	api.Put("/projects/:id", middleware.JWTMiddleware(), projectController.UpdateProject)
	api.Delete("/projects/:id", middleware.JWTMiddleware(), projectController.DeleteProject)

	api.Get("/tasks", middleware.JWTMiddleware(), taskController.GetTasks)
	api.Post("/task", middleware.JWTMiddleware(), taskController.CreateTasks)
	api.Put("/task/:id", middleware.JWTMiddleware(), taskController.UpdateTasks)
	api.Delete("/tasks/:id", middleware.JWTMiddleware(), taskController.DeleteTasks)

	if err := app.Listen(":3002"); err != nil {
		panic(err)
	}
}
