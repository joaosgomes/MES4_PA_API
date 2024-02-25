package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/gofiber/template/html/v2"
)

var engine = html.New("./views", ".html")

var FiberConfig = fiber.Config{
	ServerHeader: "MES4 - PAI", // add custom server header
	AppName:      "MES4 - PAI",
	Views:        engine,
}

var CorsConfig = cors.Config{
	Next:             nil,
	AllowOriginsFunc: nil,
	AllowOrigins:     "*",
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
	}, ","),
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           0,
}

var SwaggerConfig = swagger.Config{
	URL: "/docs/swagger.json",
}

var HealthcheckConfig = healthcheck.Config{
	LivenessProbe: func(c *fiber.Ctx) bool {
		return true
	},
	LivenessEndpoint: "/live",
	ReadinessProbe: func(c *fiber.Ctx) bool {
		return true
	},
	ReadinessEndpoint: "/ready",
}
