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

func TestBillService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
            "_id": 254894,
            "invoiceNumber": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postingDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "dueAmount": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "VendorId": 22,
            "LocationId": 61,
            "APAccountId": 4730
        }], "totalItems": 1}`)
	})

	bills, totalItems, _, err := client.Bill.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Bill.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*Bill{{
		ID:            Int64(254894),
		InvoiceNumber: String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostingDate:   String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		DueAmount:     Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		VendorId:      Int64(22),
		LocationId:    Int64(61),
		APAccountId:   Int64(4730),
	}}
	if !reflect.DeepEqual(bills, want) {
		t.Errorf("Bill.All returned %+v, want %+v", bills, want)
	}
}

func TestBillService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
            "_id": 254894,
            "invoiceNumber": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postingDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "dueAmount": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "VendorId": 22,
            "LocationId": 61,
            "APAccountId": 4730
        }`)
	})

	cc, _, err := client.Bill.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Bill.One returned error: %v", err)
	}

	want := &Bill{
		ID:            Int64(254894),
		InvoiceNumber: String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostingDate:   String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		DueAmount:     Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		VendorId:      Int64(22),
		LocationId:    Int64(61),
		APAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Bill.One returned %+v, want %+v", cc, want)
	}
}

func TestBillService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
            "_id": 254894,
            "invoiceNumber": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postingDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "dueAmount": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "VendorId": 22,
            "LocationId": 61,
            "APAccountId": 4730
        }`)
	})

	payload := &Bill{
		DueDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate: String("2017-04-29T00:00:00.000Z"),
		InvoiceDate: String("2017-04-29T00:00:00.000Z"),
		Amount:      Float64(1809.15),
		DueAmount:   Float64(1161.15),
		Status:      String("partiallyPaid"),
		Currency:    String("USD"),
		VendorId:    Int64(22),
		LocationId:  Int64(61),
		APAccountId: Int64(4730),
	}

	cc, _, err := client.Bill.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Bill.Create returned error: %v", err)
	}

	want := &Bill{
		ID:            Int64(254894),
		InvoiceNumber: String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostingDate:   String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		DueAmount:     Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		VendorId:      Int64(22),
		LocationId:    Int64(61),
		APAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Bill.Create returned %+v, want %+v", cc, want)
	}
}

func TestBillService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
            "_id": 254894,
            "invoiceNumber": "8393",
            "dueDate": "2017-05-30T00:00:00.000Z",
            "postingDate": "2017-04-29T00:00:00.000Z",
            "invoiceDate": "2017-04-29T00:00:00.000Z",
            "amount": 1809.15,
            "dueAmount": 1161.15,
            "status": "partiallyPaid",
            "currency": "USD",
            "VendorId": 22,
            "LocationId": 61,
            "APAccountId": 4730
        }`)
	})

	payload := &Bill{
		DueDate: String("2017-05-30T00:00:00.000Z"),
	}

	cc, _, err := client.Bill.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Bill.Update returned error: %v", err)
	}

	want := &Bill{
		ID:            Int64(254894),
		InvoiceNumber: String("8393"),
		DueDate:       String("2017-05-30T00:00:00.000Z"),
		PostingDate:   String("2017-04-29T00:00:00.000Z"),
		InvoiceDate:   String("2017-04-29T00:00:00.000Z"),
		Amount:        Float64(1809.15),
		DueAmount:     Float64(1161.15),
		Status:        String("partiallyPaid"),
		Currency:      String("USD"),
		VendorId:      Int64(22),
		LocationId:    Int64(61),
		APAccountId:   Int64(4730),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Bill.Update returned %+v, want %+v", cc, want)
	}
}

func TestBillService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Bill.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Bill.Delete returned error: %v", err)
	}
}

func TestBillService_void(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills/void/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Bill.Void(context.Background(), 1)
	if err != nil {
		t.Errorf("Bill.Void returned error: %v", err)
	}
}

func TestBillService_approve(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bills/approve/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Bill.Approve(context.Background(), 1)
	if err != nil {
		t.Errorf("Bill.Approve returned error: %v", err)
	}
}
