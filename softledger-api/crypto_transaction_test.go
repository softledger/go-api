package softledger

import (
	"context"
	//"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
	//"strings"
	"testing"
)

func TestCryptoTransactionService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":"1","rQty":1,"rPrice":10}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.CryptoTransaction.All(context.Background(), nil)
	if err != nil {
		t.Errorf("CryptoTransaction.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*CryptoTransaction{{
		ID:     Int64(1),
		RQty:   Float64(1),
		RPrice: Float64(10),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("CryptoTransaction.All returned %+v, want %+v", ccs, want)
	}

}

func TestCryptoTransactionService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":"1","rQty":1,"rPrice":10}`)
	})

	cc, _, err := client.CryptoTransaction.One(context.Background(), 1)
	if err != nil {
		t.Errorf("CryptoTransaction.One returned error: %v", err)
	}

	want := &CryptoTransaction{
		ID:     Int64(1),
		RQty:   Float64(1),
		RPrice: Float64(10),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CryptoTransaction.One returned %+v, want %+v", cc, want)
	}

}

func TestCryptoTransactionService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":"1","rQty":1,"rPrice":10}`)
	})

	payload := &CryptoTransaction{
		ID:     Int64(1),
		RQty:   Float64(1),
		RPrice: Float64(10),
	}

	cc, _, err := client.CryptoTransaction.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("CryptoTransaction.Create returned error: %v", err)
	}

	want := &CryptoTransaction{
		ID:     Int64(1),
		RQty:   Float64(1),
		RPrice: Float64(10),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CryptoTransaction.Create returned %+v, want %+v", cc, want)
	}
}

func TestCryptoTransactionService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":"1","rQty":1,"rPrice":10}`)
	})

	payload := &CryptoTransaction{
		RQty: Float64(1),
	}

	cc, _, err := client.CryptoTransaction.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("CryptoTransaction.Update returned error: %v", err)
	}

	want := &CryptoTransaction{
		ID:     Int64(1),
		RQty:   Float64(1),
		RPrice: Float64(10),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CryptoTransaction.Update returned %+v, want %+v", cc, want)
	}
}

func TestCryptoTransactionService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.CryptoTransaction.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("CryptoTransaction.Delete returned error: %v", err)
	}
}

func TestCryptoTransactionService_void(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/crypto/transactions/lock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	d := time.Now()

	_, err := client.CryptoTransaction.Lock(context.Background(), &CryptoTransaction{
		Date: &d,
	})
	if err != nil {
		t.Errorf("CryptoTransaction.Delete returned error: %v", err)
	}
}
