package softledger

import (
	"context"
	//"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	//"strings"
	"testing"
)

func TestCoinService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/coins", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
            "_id": "3fc2dda9-b497-4764-9f92-2bcb531ad1dc",
            "name": "0x",
            "symbol": "ZRX",
        }], "totalItems": 1}`)
	})

	coins, totalItems, _, err := client.Coin.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Coin.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Coin{{
		ID:     String("3fc2dda9-b497-4764-9f92-2bcb531ad1dc"),
		Name:   String("0x"),
		Symbol: String("ZRX"),
	}}
	if !reflect.DeepEqual(coins, want) {
		t.Errorf("Coin.All returned %+v, want %+v", coins, want)
	}
}

func TestCoinService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/coins/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
            "_id": "3fc2dda9-b497-4764-9f92-2bcb531ad1dc",
            "name": "0x",
            "symbol": "ZRX",
        }`)
	})

	cc, _, err := client.Coin.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Coin.One returned error: %v", err)
	}

	want := &Coin{
		ID:     String("3fc2dda9-b497-4764-9f92-2bcb531ad1dc"),
		Name:   String("0x"),
		Symbol: String("ZRX"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Coin.One returned %+v, want %+v", cc, want)
	}
}

func TestCoinService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/coins", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
            "_id": "3fc2dda9-b497-4764-9f92-2bcb531ad1dc",
            "name": "0x",
            "symbol": "ZRX",
        }`)
	})

	payload := &Coin{
		Name:   String("0x"),
		Symbol: String("ZRX"),
	}

	cc, _, err := client.Coin.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Coin.Create returned error: %v", err)
	}

	want := &Coin{
		ID:     String("3fc2dda9-b497-4764-9f92-2bcb531ad1dc"),
		Name:   String("0x"),
		Symbol: String("ZRX"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Coin.Create returned %+v, want %+v", cc, want)
	}
}

func TestCoinService_hide(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/coins/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Coin.Hide(context.Background(), 1)
	if err != nil {
		t.Errorf("Coin.Hide returned error: %v", err)
	}
}
