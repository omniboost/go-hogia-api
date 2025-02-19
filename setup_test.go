package hogia_api_test

import (
	"context"
	"os"
	"testing"

	hogia_api "github.com/omniboost/go-hogia-api"
)

var (
	client *hogia_api.Client
)

func TestMain(m *testing.M) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	// Default oausth2 flow
	oauthConfig := hogia_api.NewOauth2ClientCredentialsConfig()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = hogia_api.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
