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

func TestLedgerAccountService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ledger_accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
			"_id":1,
			"name":"1",
			"number":"one",
			"description": "1"
		}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.LedgerAccount.All(context.Background(), nil)
	if err != nil {
		t.Errorf("LedgerAccount.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong description of ledger_accounts, want %v, got %v", 1, totalItems)
	}

	want := []*LedgerAccount{{
		ID:          Int64(1),
		Name:        String("1"),
		Number:      String("one"),
		Description: String("1"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("LedgerAccount.All returned %+v, want %+v", ccs, want)
	}

}

func TestLedgerAccountService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ledger_accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"number":"one",
			"description": "1"
		}`)
	})

	cc, _, err := client.LedgerAccount.One(context.Background(), 1)
	if err != nil {
		t.Errorf("LedgerAccount.One returned error: %v", err)
	}

	want := &LedgerAccount{
		ID:          Int64(1),
		Name:        String("1"),
		Number:      String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("LedgerAccount.One returned %+v, want %+v", cc, want)
	}

}

func TestLedgerAccountService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ledger_accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"number":"one",
			"description": "1"
		}`)
	})

	payload := &LedgerAccount{
		Name:   String("1"),
		Number: String("one"),
	}

	cc, _, err := client.LedgerAccount.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("LedgerAccount.Create returned error: %v", err)
	}

	want := &LedgerAccount{
		ID:          Int64(1),
		Name:        String("1"),
		Number:      String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("LedgerAccount.Create returned %+v, want %+v", cc, want)
	}
}

func TestLedgerAccountService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ledger_accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"number":"one",
			"description": "1"
		}`)
	})

	payload := &LedgerAccount{
		Name: String("1"),
	}

	cc, _, err := client.LedgerAccount.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("LedgerAccount.Update returned error: %v", err)
	}

	want := &LedgerAccount{
		ID:          Int64(1),
		Name:        String("1"),
		Number:      String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("LedgerAccount.Update returned %+v, want %+v", cc, want)
	}
}

func TestLedgerAccountService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ledger_accounts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.LedgerAccount.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("LedgerAccount.Delete returned error: %v", err)
	}
}
