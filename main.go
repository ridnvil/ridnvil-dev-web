package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"ridnvil-dev/pkg/controllers"
	"ridnvil-dev/pkg/database"
	"ridnvil-dev/pkg/models"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "ridnvil.dev"})
	app.Use(logger.New())

	app.Static("/", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/login", "ridnvil/build", fiber.Static{Browse: false})
	app.Static("/profile", "ridnvil/build", fiber.Static{Browse: false})

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
		&models.Experinces{},
		&models.SocialNetwork{},
		&models.Educations{},
	); err != nil {
		panic(err)
	}

	api := app.Group("/api")

	// Testing Push
	api.Post("/login", controllers.Login)

	api.Get("/welcome", controllers.Welcome)
	api.Post("/client", controllers.CreateClient)
	api.Get("/client", controllers.GetListClient)

	api.Get("/profile", controllers.GetProfiles)

	api.Get("/stock", controllers.GetStock)

	if err := app.Listen(":3001"); err != nil {
		panic(err)
	}
}
