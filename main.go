package main

import (
	"api/config"
	controller "api/controllers"
	"api/repository"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	// Middleware
	app.Use(logger.New())

	// Cors
	app.Use(cors.New(config.CorsConfig))

	// Swagger documentation
	app.Get("/swagger/*", swagger.New(config.SwaggerConfig))
	app.Static("/docs", "./docs")

	// Healthcheck
	app.Use(healthcheck.New(config.HealthcheckConfig))

	// Redirect root to Swagger UI
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger")
	})

	// Initialize event repository
	eventRepository := repository.NewEventRepository()
	// Initialize image repository
	imageRepository := repository.NewImageRepository()

	// Initialize event controller
	eventController := controller.NewEventController(eventRepository)
	// Initialize image controller
	imageController := controller.NewImageController(imageRepository)
	// Initialize Template controller
	eventTemplateController := controller.EventTemplateRepository(eventRepository)

	// Metrics
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	// Routes
	app.Get("/event/:id", eventController.GetEvent())
	app.Get("/event", eventController.GetEvents())
	app.Post("/event", eventController.PostEvent())
	app.Put("/event/:id", eventController.PutEvent())
	app.Delete("/event/:id", eventController.DeleteEvent())
	app.Get("/ws/:id", websocket.New(eventController.WSHandler))

	app.Get("/image/:id", imageController.GetImage())
	app.Get("/image", imageController.GetImages())
	app.Post("/image", imageController.PostImage())
	app.Delete("/image/:id", imageController.DeleteImage())

	// Serve static files
	app.Static("/", "./public")

	//Template (HTML)
	app.Get("/events", eventTemplateController.GetEventsTemplate)
	app.Get("/events/:id", eventTemplateController.GetEventTemplate)

	// Listen on port 5000
	app.Listen(":5000")
}
