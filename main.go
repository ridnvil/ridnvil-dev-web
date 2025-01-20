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

	app.Static("/", "ridnvil/build", fiber.Static{
		Browse: false,
	})

	dash := app.Group("/admin")
	app.Static("/assets", "dashboard/dist/assets")
	app.Static("/", "dashboard/vite.svg")
	dash.Static("/", "dashboard/dist")

	if errdbcheck := database.AutoCreateDatabase(); errdbcheck != nil {
		log.Println(errdbcheck)
	}

	db, err := database.OpenConnectionSQLite()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.IPInfo{}, &models.Experinces{}, &models.Profile{}, &models.Educations{}, &models.SocialNetwork{}); err != nil {
		panic(err)
	}

	api := app.Group("/api")

	api.Get("/welcome", controllers.Welcome)
	api.Post("/client", controllers.CreateClient)
	api.Get("/client", controllers.GetListClient)

	api.Get("/stock", controllers.GetStock)

	if err := app.Listen(":3001"); err != nil {
		panic(err)
	}
}
