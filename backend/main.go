package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maxkruse/osu-popularityperformance/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "ppV2.5",
		StrictRouting: true,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	// api grouping
	app.Get("/api/oauth/start", routes.OAuthStart)
	app.Get("/api/oauth/callback", routes.OAuthCallback)

	fmt.Errorf("%s", app.Listen(":5000"))
}
