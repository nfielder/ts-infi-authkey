package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nfielder/ts-infi-authkey/internal/client"
	"golang.org/x/oauth2/clientcredentials"
)

type CmdOpts struct {
	Reusable  bool
	Ephemeral bool
	Preauth   bool
	Tags      string
	Expiry    time.Duration
}

const maxDuration time.Duration = 90 * 24 * time.Hour

func Run(opts CmdOpts) {
	clientID := os.Getenv("TS_API_CLIENT_ID")
	clientSecret := os.Getenv("TS_API_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("TS_API_CLIENT_ID and TS_API_CLIENT_SECRET must be set")
	}

	if opts.Tags == "" {
		log.Fatal("at least one tag must be specified")
	}

	if opts.Expiry > maxDuration {
		log.Fatal("max expiry time must be less than 90 days")
	}

	baseURL := os.Getenv("TS_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.tailscale.com"
	}

	credentials := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     baseURL + "/api/v2/oauth/token",
		Scopes:       []string{"device"},
	}

	ctx := context.Background()
	// Create client with no specific Tailnet
	tsClient := client.NewClient("-")
	tsClient.HTTPClient = credentials.Client(ctx)
	tsClient.BaseURL = baseURL

	caps := client.KeyCapabilities{
		Devices: client.KeyDeviceCapabilities{
			Create: client.KeyDeviceCreateCapabilities{
				Reusable:      opts.Reusable,
				Ephemeral:     opts.Ephemeral,
				Preauthorized: opts.Preauth,
				Tags:          strings.Split(opts.Tags, ","),
			},
		},
	}

	authkey, _, err := tsClient.CreateKeyWithExpiry(ctx, caps, opts.Expiry)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(authkey)
}
