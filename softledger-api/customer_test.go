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

func TestCustomerService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"name":"one", "Addresses": [{"_id": 1}]}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Customer.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Customer.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Customer{{
		ID:   Int64(1),
		Name: String("one"),
		Addresses: &Addresses{{
			ID: Int64(1),
		}},
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Customer.All returned %+v, want %+v", ccs, want)
	}

}

func TestCustomerService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	cc, _, err := client.Customer.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Customer.One returned error: %v", err)
	}

	want := &Customer{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Customer.One returned %+v, want %+v", cc, want)
	}

}

func TestCustomerService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	payload := &Customer{
		Name: String("one"),
	}

	cc, _, err := client.Customer.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Customer.Create returned error: %v", err)
	}

	want := &Customer{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Customer.Create returned %+v, want %+v", cc, want)
	}
}

func TestCustomerService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":1,"name":"two"}`)
	})

	payload := &Customer{
		Name: String("two"),
	}

	cc, _, err := client.Customer.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Customer.Update returned error: %v", err)
	}

	want := &Customer{
		ID:   Int64(1),
		Name: String("two"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Customer.Update returned %+v, want %+v", cc, want)
	}
}

func TestCustomerService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/customers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Customer.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Customer.Delete returned error: %v", err)
	}
}
