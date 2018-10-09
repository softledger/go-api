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

func TestWarehouseService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Warehouse.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Warehouse.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Warehouse{{
		ID:   Int64(1),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Warehouse.All returned %+v, want %+v", ccs, want)
	}

}

func TestWarehouseService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouses/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	cc, _, err := client.Warehouse.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Warehouse.One returned error: %v", err)
	}

	want := &Warehouse{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Warehouse.One returned %+v, want %+v", cc, want)
	}

}

func TestWarehouseService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	payload := &Warehouse{
		Name: String("one"),
	}

	cc, _, err := client.Warehouse.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Warehouse.Create returned error: %v", err)
	}

	want := &Warehouse{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Warehouse.Create returned %+v, want %+v", cc, want)
	}
}

func TestWarehouseService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouses/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":1,"name":"two"}`)
	})

	payload := &Warehouse{
		Name: String("two"),
	}

	cc, _, err := client.Warehouse.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Warehouse.Update returned error: %v", err)
	}

	want := &Warehouse{
		ID:   Int64(1),
		Name: String("two"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Warehouse.Update returned %+v, want %+v", cc, want)
	}
}

func TestWarehouseService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouses/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Warehouse.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Warehouse.Delete returned error: %v", err)
	}
}
