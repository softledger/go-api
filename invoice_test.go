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

func TestInvoiceService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
            "_id": 254894,
            "number": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postedDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "amountPayable": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "AgentId": 22,
            "LocationId": 61,
            "ARAccountId": 4730
        }], "totalItems": 1}`)
	})

	invoices, totalItems, _, err := client.Invoice.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Invoice.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Invoice{{
		ID:            Int64(254894),
		Number:        String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostedDate:    String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		AmountPayable: Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		AgentId:       Int64(22),
		LocationId:    Int64(61),
		ARAccountId:   Int64(4730),
	}}
	if !reflect.DeepEqual(invoices, want) {
		t.Errorf("Invoice.All returned %+v, want %+v", invoices, want)
	}
}

func TestInvoiceService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
            "_id": 254894,
            "number": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postedDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "amountPayable": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "AgentId": 22,
            "LocationId": 61,
            "ARAccountId": 4730
        }`)
	})

	cc, _, err := client.Invoice.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Invoice.One returned error: %v", err)
	}

	want := &Invoice{
		ID:            Int64(254894),
		Number:        String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostedDate:    String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		AmountPayable: Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		AgentId:       Int64(22),
		LocationId:    Int64(61),
		ARAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Invoice.One returned %+v, want %+v", cc, want)
	}
}

func TestInvoiceService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
            "_id": 254894,
            "number": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postedDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "amountPayable": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "AgentId": 22,
            "LocationId": 61,
            "ARAccountId": 4730
        }`)
	})

	payload := &Invoice{
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostedDate:    String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		AmountPayable: Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		AgentId:       Int64(22),
		LocationId:    Int64(61),
		ARAccountId:   Int64(4730),
	}

	cc, _, err := client.Invoice.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Invoice.Create returned error: %v", err)
	}

	want := &Invoice{
		ID:            Int64(254894),
		Number:        String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostedDate:    String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		AmountPayable: Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		AgentId:       Int64(22),
		LocationId:    Int64(61),
		ARAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Invoice.Create returned %+v, want %+v", cc, want)
	}
}

func TestInvoiceService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
            "_id": 254894,
            "number": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postedDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "amountPayable": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "AgentId": 22,
            "LocationId": 61,
            "ARAccountId": 4730
        }`)
	})

	payload := &Invoice{
		DueDate: String("2017-05-30T00:00:00.000Z"),
	}

	cc, _, err := client.Invoice.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Invoice.Update returned error: %v", err)
	}

	want := &Invoice{
		ID:            Int64(254894),
		Number:        String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostedDate:    String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		AmountPayable: Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		AgentId:       Int64(22),
		LocationId:    Int64(61),
		ARAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Invoice.Update returned %+v, want %+v", cc, want)
	}
}

func TestInvoiceService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Invoice.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Invoice.Delete returned error: %v", err)
	}
}

func TestInvoiceService_issue(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices/1/issue", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Invoice.Issue(context.Background(), 1)
	if err != nil {
		t.Errorf("Invoice.Issue returned error: %v", err)
	}
}

func TestInvoiceService_void(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/invoices/1/void", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Invoice.Void(context.Background(), 1)
	if err != nil {
		t.Errorf("Invoice.Void returned error: %v", err)
	}
}
