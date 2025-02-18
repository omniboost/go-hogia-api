package hogia_v2_test

import (
	"context"
	"os"
	"testing"

	hogia_v2 "github.com/omniboost/go-hogia-v2"
)

var (
	client *hogia_v2.Client
)

func TestMain(m *testing.M) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	// Default oausth2 flow
	oauthConfig := hogia_v2.NewOauth2ClientCredentialsConfig()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = hogia_v2.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
