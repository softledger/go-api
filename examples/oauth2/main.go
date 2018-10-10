package main

import (
	"context"
	"fmt"
	"github.com/softledger/go-api/softledger-api"
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
			"tenantUUID": {"71618775-bc61-41c7-ad0b-0c19a53fa3c4"},
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
