package routes

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxkruse/osu-popularityperformance/database"
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

	// make db session
	databaseSession := database.GetSession().Debug()

	// Search for this user
	user := models.User{}

	databaseSession.Preload("Sessions").First(&user, "bancho_id = ?", userInfo.UserId)

	// make a new sessionToken
	sessToken := generateState(16)

	user.BanchoId = userInfo.UserId
	user.Username = userInfo.Username
	user.AvatarUri = userInfo.AvatarURL
	user.Sessions = append(user.Sessions, models.Session{
		SessionToken: sessToken,
	})

	// delete all tokens that were granted to the user. Dont just set the deleted_at, actually delete it
	databaseSession.Unscoped().Delete(&models.Token{}, "user_id = ?", user.ID)

	user.OAuthTokens = &models.Token{
		Token: *tok,
	}

	databaseSession.Save(&user)

	// - redirect to frontend, giving the SessionToken as a HTTP-Param

	// Get redirect URL for the frontend:
	frontendUri := os.Getenv("FRONTEND_REDIRECT_URI")
	return c.Redirect(frontendUri + "?session=" + sessToken)
}
