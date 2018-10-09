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

	mux.HandleFunc("/bills", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{"_id":1,"id":"1","name":"one"}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.CostCenter.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Bill.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Bill{{
		ID:   Int64(1),
		Id:   String("1"),
		Name: String("one"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("CostCenter.All returned %+v, want %+v", ccs, want)
	}

}
