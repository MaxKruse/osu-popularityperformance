package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maxkruse/osu-popularityperformance/database"
	"github.com/maxkruse/osu-popularityperformance/routes"
)

func main() {
	// seed the rng
	rand.Seed(time.Now().UnixNano())

	app := fiber.New(fiber.Config{
		AppName:       "ppV2.5",
		StrictRouting: true,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	if _, exists := os.LookupEnv("DEBUG"); exists {
		log.Println("Debug mode enabled")
		app.Use(cors.New())
	}

	// start database connection
	database.Init()

	// api grouping
	app.Get("/api/oauth/start", routes.OAuthStart)
	app.Get("/api/oauth/callback", routes.OAuthCallback)

	log.Fatalf("%s", app.Listen(":5000"))
}
