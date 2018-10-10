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

func TestVendorService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/vendors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Vendor.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Vendor.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Vendor{{
		ID:   Int64(1),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Vendor.All returned %+v, want %+v", ccs, want)
	}

}

func TestVendorService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/vendors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	cc, _, err := client.Vendor.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Vendor.One returned error: %v", err)
	}

	want := &Vendor{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Vendor.One returned %+v, want %+v", cc, want)
	}

}

func TestVendorService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/vendors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	payload := &Vendor{
		Name: String("one"),
	}

	cc, _, err := client.Vendor.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Vendor.Create returned error: %v", err)
	}

	want := &Vendor{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Vendor.Create returned %+v, want %+v", cc, want)
	}
}

func TestVendorService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/vendors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":1,"name":"two"}`)
	})

	payload := &Vendor{
		Name: String("two"),
	}

	cc, _, err := client.Vendor.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Vendor.Update returned error: %v", err)
	}

	want := &Vendor{
		ID:   Int64(1),
		Name: String("two"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Vendor.Update returned %+v, want %+v", cc, want)
	}
}

func TestVendorService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/vendors/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Vendor.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Vendor.Delete returned error: %v", err)
	}
}
