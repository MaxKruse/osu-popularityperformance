package routes

import (
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

var (
	conf *oauth2.Config
)

func generateState() string {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, len(alphabet))
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}

func OAuthStart(c *fiber.Ctx) error {
	// get needed oAuth params from environment
	clientID := os.Getenv("OSU_BANCHO_CLIENT_ID")
	clientSecret := os.Getenv("OSU_BANCHO_CLIENT_SECRET")
	redirectUri := os.Getenv("OSU_BANCHO_CLIENT_CALLBACK")

	// make oauth client
	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"public"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://osu.ppy.sh/oauth/authorize",
			TokenURL: "https://osu.ppy.sh/oauth/token",
		},
		RedirectURL: redirectUri,
	}

	// generate random state for later MITM prevention
	state := generateState()

	// set state for this session
	c.Cookie(&fiber.Cookie{
		Name:  "state",
		Value: state,
	})

	url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

	// redirect to osu! oauth
	return c.Redirect(url)
}
