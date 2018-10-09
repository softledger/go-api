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

func TestLocationService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
			"_id":1,
			"name":"1",
			"id":"one",
			"description": "1"
		}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Location.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Location.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong description of locations, want %v, got %v", 1, totalItems)
	}

	want := []*Location{{
		ID:          Int64(1),
		Name:        String("1"),
		Id:          String("one"),
		Description: String("1"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Location.All returned %+v, want %+v", ccs, want)
	}

}

func TestLocationService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"id":"one",
			"description": "1"
		}`)
	})

	cc, _, err := client.Location.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Location.One returned error: %v", err)
	}

	want := &Location{
		ID:          Int64(1),
		Name:        String("1"),
		Id:          String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Location.One returned %+v, want %+v", cc, want)
	}

}

func TestLocationService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"id":"one",
			"description": "1"
		}`)
	})

	payload := &Location{
		Name: String("1"),
		Id:   String("one"),
	}

	cc, _, err := client.Location.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Location.Create returned error: %v", err)
	}

	want := &Location{
		ID:          Int64(1),
		Name:        String("1"),
		Id:          String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Location.Create returned %+v, want %+v", cc, want)
	}
}

func TestLocationService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"_id":1,
			"name":"1",
			"id":"one",
			"description": "1"
		}`)
	})

	payload := &Location{
		Name: String("1"),
	}

	cc, _, err := client.Location.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Location.Update returned error: %v", err)
	}

	want := &Location{
		ID:          Int64(1),
		Name:        String("1"),
		Id:          String("one"),
		Description: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Location.Update returned %+v, want %+v", cc, want)
	}
}

func TestLocationService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/locations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Location.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Location.Delete returned error: %v", err)
	}
}
