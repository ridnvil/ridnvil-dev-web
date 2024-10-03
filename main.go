package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"ridnvil-dev/pkg/controllers"
	"ridnvil-dev/pkg/database"
	"ridnvil-dev/pkg/models"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "ridnvil.dev"})
	app.Use(logger.New())

	app.Static("/", "ridnvil/build", fiber.Static{
		Browse: false,
	})

	db, err := database.OpenConnectionSQLite()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.IPInfo{}); err != nil {
		panic(err)
	}

	api := app.Group("/api")

	api.Get("/welcome", controllers.Welcome)
	api.Post("/client", controllers.CreateClient)

	if err := app.Listen(":3001"); err != nil {
		panic(err)
	}
}
