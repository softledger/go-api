# go-api
GO api client

This is still very much an alpha pkg.

Endpoints not implemented:
- Kit
- Location
- PurchaseOrders
- Transactions
- SalesOrders
- Stock
- Vendor
- Wallet
- Warehouse


Example uses with oauth2

```
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
		ClientID:     "MY_CLIENT_ID",
		ClientSecret: "MY_CLIENT_SECRET",
		TokenURL:     "https://softledger.auth0.com/oauth/token",
		EndpointParams: url.Values{
			"audience":   {"https://sl-prod.softledger.com/api"},
			"tenantUUID": {"MY_TENANTID"},
		},
	}

	ctx := context.Background()

	client := softledger.NewClient(config.Client(ctx))

	_, totalItems, _, err := client.LedgerAccount.All(ctx, nil)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("totalItems", totalItems)

}
```