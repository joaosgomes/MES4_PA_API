package main

import (
	"api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/swagger"
)

func main() {

	// Start fiber app
	app := fiber.New(config.FiberConfig)

	//Cors
	app.Use(cors.New(config.CorsConfig))

	app.Get("/swagger/*", swagger.New(config.SwaggerConfig))
	app.Static("/docs", "./docs")

	app.Use(healthcheck.New(config.HealthcheckConfig))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger")
	})

	app.Listen(":5000")
}
