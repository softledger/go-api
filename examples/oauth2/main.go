package main

import (
	"context"
	"fmt"
	"github.com/softledger/go-api/softledger-api"
	"golang.org/x/oauth2/clientcredentials"
	"net/url"
)

const (
	CLIENT_ID     = "client_id"
	CLIENT_SECRET = "client_secret"
	TENANTID      = "tenantId"
)

func main() {

	config := clientcredentials.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		TokenURL:     "https://softledger.auth0.com/oauth/token",
		EndpointParams: url.Values{
			"audience":   {"https://sl-dev.softledger.com/api"},
			"tenantUUID": {TENANTID},
		},
	}

	ctx := context.Background()

	client := softledger.NewClient(config.Client(ctx))
	client.BaseURL, _ = url.Parse("https://dev-api.softledger.com/")

	//fmt.Println("client", client)

	_, totalItems, _, err := client.LedgerAccount.All(ctx, nil)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("totalItems", totalItems)

}
