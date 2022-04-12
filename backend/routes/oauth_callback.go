package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/maxkruse/osu-popularityperformance/models"
)

func OAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")
	cookieState := c.Cookies("state")

	// invalidate state cookie
	c.ClearCookie("state")

	// check if state is valid
	if cookieState != state {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid state",
		})
	}

	tok, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// we got a working token now, get some info on the user

	client := conf.Client(c.Context(), tok)
	userResp, err := client.Get("https://osu.ppy.sh/api/v2/me")
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse user info
	var userInfo models.BanchoUser
	json.NewDecoder(userResp.Body).Decode(&userInfo)

	// TODO:
	// - check if user exists
	// - add session to user's sessions
	// - save user to db
	// - redirect to frontend, giving the SessionToken as a HTTP-Param

	// for now, just show what userdata we got from bancho
	return c.JSON(userInfo)
}
