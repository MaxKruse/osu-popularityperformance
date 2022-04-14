package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/maxkruse/osu-popularityperformance/database"
	"github.com/maxkruse/osu-popularityperformance/models"
)

func GetSelf(c *fiber.Ctx) error {
	// get authorization header, removing the "Bearer " prefix
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.Contains(authHeader, "Bearer ") || authHeader[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{
			"error": "No authorization header found",
		})
	}

	// get user from database
	user := models.User{}
	sess := models.Session{
		SessionToken: authHeader[7:],
	}

	dbSess := database.GetSession()

	if err := dbSess.First(&sess, sess).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid session header",
		})
	}

	if err := dbSess.Preload("Sessions").First(&user, sess.UserId).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "no user found",
		})
	}

	return c.JSON(user)
}
