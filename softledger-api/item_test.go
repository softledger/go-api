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

func TestItemService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
			"_id":1,
			"sku":"1",
			"name":"one",
			"number": "1"
		}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Item.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Item.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Item{{
		ID:     Int64(1),
		Sku:    String("1"),
		Name:   String("one"),
		Number: String("1"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Item.All returned %+v, want %+v", ccs, want)
	}

}

func TestItemService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"_id":1,
			"sku":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	cc, _, err := client.Item.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Item.One returned error: %v", err)
	}

	want := &Item{
		ID:     Int64(1),
		Sku:    String("1"),
		Name:   String("one"),
		Number: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Item.One returned %+v, want %+v", cc, want)
	}

}

func TestItemService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
			"_id":1,
			"sku":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	payload := &Item{
		Sku:  String("1"),
		Name: String("one"),
	}

	cc, _, err := client.Item.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Item.Create returned error: %v", err)
	}

	want := &Item{
		ID:     Int64(1),
		Sku:    String("1"),
		Name:   String("one"),
		Number: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Item.Create returned %+v, want %+v", cc, want)
	}
}

func TestItemService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"_id":1,
			"sku":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	payload := &Item{
		Sku: String("1"),
	}

	cc, _, err := client.Item.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Item.Update returned error: %v", err)
	}

	want := &Item{
		ID:     Int64(1),
		Sku:    String("1"),
		Name:   String("one"),
		Number: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Item.Update returned %+v, want %+v", cc, want)
	}
}

func TestItemService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Item.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Item.Delete returned error: %v", err)
	}
}
