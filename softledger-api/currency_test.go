package softledger

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStatusService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/currency", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `[    {
        "fraction": 2,
        "rounding_method": 1,
        "code": "AFN",
        "name": "Afghani",
        "symbol": "؋"
    },
    {
        "fraction": 2,
        "rounding_method": 1,
        "code": "EUR",
        "name": "Euro",
        "symbol": "€"
    }]`)
	})

	cc, _, err := client.Currency.All(context.Background())
	if err != nil {
		t.Errorf("Currency.All returned error: %v", err)
	}

	want := []*Currency{{
		Fraction:       Int64(2),
		RoundingMethod: Int64(1),
		Code:           String("AFN"),
		Name:           String("Afghani"),
		Symbol:         String("؋"),
	}, {
		Fraction:       Int64(2),
		RoundingMethod: Int64(1),
		Code:           String("EUR"),
		Name:           String("Euro"),
		Symbol:         String("€"),
	}}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Currency.All returned %+v, want %+v", cc, want)
	}

}
