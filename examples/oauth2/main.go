package main

import (
	"context"
	"fmt"
	"github.com/softledger/go-api/softledger"
	"golang.org/x/oauth2/clientcredentials"
	"net/url"
)

func main() {

	config := clientcredentials.Config{
		ClientID:     "Ng5ho5FmupnFlpeZ0HnZukMPBDQxO0HK",
		ClientSecret: "R5FBmB91XgYvAGeS39w3nPhCoNCGUqRSLt3ri8IyCaxHjMKflOrFNvIPkp7MuK9c",
		TokenURL:     "https://softledger.auth0.com/oauth/token",
		EndpointParams: url.Values{
			"audience":   {"https://sl-dev.softledger.com/api"},
			"tenantUUID": {"6f825930-509e-496c-8778-54b1eeed4033"},
		},
	}

	client := softledger.NewClient(config.Client)
	ctx := context.Background()

	accounts, totalItems, err := client.LedgerAccounts.All(ctx)

	fmt.Println("totalItems", totalItems)

}
