package main

import (
	"api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/swagger"
)

// @title MES4_PA_API
// @version 1.0
// @description Description MES4_PA_API
// @description Description MES4_PA_API GO
// @description Description MES4_PA_API GO Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name João Da Silva Gomes
// @contact.url https://joaosilvagomes.pt
// @contact.email joao.s.gomes@outlook.pt
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
// @externalDocs.url https://swagger.io/resources/open-api/
// @externalDocs.description OpenAPI
// @host localhost:5000
// @BasePath /
// @tag.name MES4_PA_API
// @tag.description MES4_PA_API
// @tag.docs.url https://github.com/joaosgomes/MES4_PA_API
// @tag.docs.description https://github.com/joaosgomes/MES4_PA_API
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
