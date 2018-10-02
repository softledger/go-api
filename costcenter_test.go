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

func TestCostCenterService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cost_centers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"id":"1","name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.CostCenter.All(context.Background(), nil)
	if err != nil {
		t.Errorf("CostCenter.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*CostCenter{{
		ID:   Int64(1),
		Id:   String("1"),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("CostCenter.All returned %+v, want %+v", ccs, want)
	}

}

func TestCostCenterService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cost_centers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"_id":1,"id":"1","name":"one"}`)
	})

	cc, _, err := client.CostCenter.One(context.Background(), 1)
	if err != nil {
		t.Errorf("CostCenter.One returned error: %v", err)
	}

	want := &CostCenter{
		ID:   Int64(1),
		Id:   String("1"),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CostCenter.One returned %+v, want %+v", cc, want)
	}

}

func TestCostCenterService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cost_centers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{"_id":1,"id":"1","name":"one"}`)
	})

	payload := &CostCenter{
		Id:   String("1"),
		Name: String("one"),
	}

	cc, _, err := client.CostCenter.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("CostCenter.Create returned error: %v", err)
	}

	want := &CostCenter{
		ID:   Int64(1),
		Id:   String("1"),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CostCenter.Create returned %+v, want %+v", cc, want)
	}
}

func TestCostCenterService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cost_centers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"_id":1,"id":"2","name":"one"}`)
	})

	payload := &CostCenter{
		Id: String("2"),
	}

	cc, _, err := client.CostCenter.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("CostCenter.Update returned error: %v", err)
	}

	want := &CostCenter{
		ID:   Int64(1),
		Id:   String("2"),
		Name: String("one"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CostCenter.Update returned %+v, want %+v", cc, want)
	}
}

func TestCostCenterService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cost_centers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.CostCenter.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("CostCenter.Delete returned error: %v", err)
	}
}
