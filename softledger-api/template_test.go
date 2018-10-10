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

func TestTemplateService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Template.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Template.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Template{{
		ID:   Int64(1),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Template.All returned %+v, want %+v", ccs, want)
	}

}

func TestTemplateService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	cc, _, err := client.Template.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Template.One returned error: %v", err)
	}

	want := &Template{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Template.One returned %+v, want %+v", cc, want)
	}

}

func TestTemplateService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":1,"name":"one"}`)
	})

	payload := &Template{
		Name: String("one"),
	}

	cc, _, err := client.Template.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Template.Create returned error: %v", err)
	}

	want := &Template{
		ID:   Int64(1),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Template.Create returned %+v, want %+v", cc, want)
	}
}

func TestTemplateService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":1,"name":"two"}`)
	})

	payload := &Template{
		Name: String("two"),
	}

	cc, _, err := client.Template.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Template.Update returned error: %v", err)
	}

	want := &Template{
		ID:   Int64(1),
		Name: String("two"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Template.Update returned %+v, want %+v", cc, want)
	}
}

func TestTemplateService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Template.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Template.Delete returned error: %v", err)
	}
}
