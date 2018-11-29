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

func TestWalletService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":"1","name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Wallet.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Wallet.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Wallet{{
		ID:   String("1"),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Wallet.All returned %+v, want %+v", ccs, want)
	}

}

func TestWalletService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wallets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":"1","name":"one"}`)
	})

	cc, _, err := client.Wallet.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Wallet.One returned error: %v", err)
	}

	want := &Wallet{
		ID:   String("1"),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Wallet.One returned %+v, want %+v", cc, want)
	}

}

func TestWalletService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":"1","name":"one"}`)
	})

	payload := &Wallet{
		Name: String("one"),
	}

	cc, _, err := client.Wallet.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Wallet.Create returned error: %v", err)
	}

	want := &Wallet{
		ID:   String("1"),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Wallet.Create returned %+v, want %+v", cc, want)
	}
}

func TestWalletService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wallets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":"1","name":"two"}`)
	})

	payload := &Wallet{
		Name: String("two"),
	}

	cc, _, err := client.Wallet.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Wallet.Update returned error: %v", err)
	}

	want := &Wallet{
		ID:   String("1"),
		Name: String("two"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Wallet.Update returned %+v, want %+v", cc, want)
	}
}

func TestWalletService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wallets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Wallet.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Wallet.Delete returned error: %v", err)
	}
}
