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

func TestCashReceiptService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cashReceipts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
            "_id": 3,
            "number": "001",
            "type": "Check",
            "amount": "100",
            "unused": "0",
            "description": "Office Supplies",
            "currency": "USD",
            "receiveDate": "2017-07-05T19:57:49.964Z",
            "status": "created",
            "AgentId": 11,
            "LedgerAccountId": 4790,
            "LocationId": 62,
        }], "totalItems": 1}`)
	})

	cashReceipts, totalItems, _, err := client.CashReceipt.All(context.Background(), nil)
	if err != nil {
		t.Errorf("CashReceipt.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong number of items, want %v, got %v", 1, totalItems)
	}

	want := []*CashReceipt{{
		ID:              Int64(3),
		Number:          String("001"),
		ReceiveDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate:     String("2017-04-29T00:00:00.000Z"),
		Amount:          Float64(100),
		Type:            String("Check"),
		Description:     String("Office Supplies"),
		Unused:          Float64(0),
		Status:          String("created"),
		Currency:        String("USD"),
		AgentId:         Int64(11),
		LocationId:      Int64(61),
		LedgerAccountId: Int64(4790),
	}}
	if !reflect.DeepEqual(cashReceipts, want) {
		t.Errorf("CashReceipt.All returned %+v, want %+v", cashReceipts, want)
	}
}

func TestCashReceiptService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cashReceipts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
            "_id": 3,
            "number": "001",
            "type": "Check",
            "amount": "100",
            "unused": "0",
            "description": "Office Supplies",
            "currency": "USD",
            "receiveDate": "2017-07-05T19:57:49.964Z",
            "status": "created",
            "AgentId": 11,
            "LedgerAccountId": 4790,
            "LocationId": 62,
        }`)
	})

	cc, _, err := client.CashReceipt.One(context.Background(), 1)
	if err != nil {
		t.Errorf("CashReceipt.One returned error: %v", err)
	}

	want := &CashReceipt{
		ID:              Int64(3),
		Number:          String("001"),
		ReceiveDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate:     String("2017-04-29T00:00:00.000Z"),
		Amount:          Float64(100),
		Type:            String("Check"),
		Description:     String("Office Supplies"),
		Unused:          Float64(0),
		Status:          String("created"),
		Currency:        String("USD"),
		AgentId:         Int64(11),
		LocationId:      Int64(61),
		LedgerAccountId: Int64(4790),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CashReceipt.One returned %+v, want %+v", cc, want)
	}
}

func TestCashReceiptService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cashReceipts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
            "_id": 3,
            "number": "001",
            "type": "Check",
            "amount": "100",
            "unused": "0",
            "description": "Office Supplies",
            "currency": "USD",
            "receiveDate": "2017-07-05T19:57:49.964Z",
            "status": "created",
            "AgentId": 11,
            "LedgerAccountId": 4790,
            "LocationId": 62,
        }`)
	})

	payload := &CashReceipt{
		Number:          String("001"),
		ReceiveDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate:     String("2017-04-29T00:00:00.000Z"),
		Amount:          Float64(100),
		Type:            String("Check"),
		Description:     String("Office Supplies"),
		Unused:          Float64(0),
		Status:          String("created"),
		Currency:        String("USD"),
		AgentId:         Int64(11),
		LocationId:      Int64(61),
		LedgerAccountId: Int64(4790),
	}

	cc, _, err := client.CashReceipt.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("CashReceipt.Create returned error: %v", err)
	}

	want := &CashReceipt{
		ID:              Int64(3),
		Number:          String("001"),
		ReceiveDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate:     String("2017-04-29T00:00:00.000Z"),
		Amount:          Float64(100),
		Type:            String("Check"),
		Description:     String("Office Supplies"),
		Unused:          Float64(0),
		Status:          String("created"),
		Currency:        String("USD"),
		AgentId:         Int64(11),
		LocationId:      Int64(61),
		LedgerAccountId: Int64(4790),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CashReceipt.Create returned %+v, want %+v", cc, want)
	}
}

func TestCashReceiptService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cashReceipts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
            "_id": 3,
            "number": "001",
            "type": "Check",
            "amount": "100",
            "unused": "0",
            "description": "Office Supplies",
            "currency": "USD",
            "receiveDate": "2017-07-05T19:57:49.964Z",
            "status": "created",
            "AgentId": 11,
            "LedgerAccountId": 4790,
            "LocationId": 62,
        }`)
	})

	payload := &CashReceipt{
		Number: String("001"),
	}

	cc, _, err := client.CashReceipt.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("CashReceipt.Update returned error: %v", err)
	}

	want := &CashReceipt{
		ID:              Int64(3),
		Number:          String("001"),
		ReceiveDate:     String("2017-05-30T00:00:00.000Z"),
		PostingDate:     String("2017-04-29T00:00:00.000Z"),
		Amount:          Float64(100),
		Type:            String("Check"),
		Description:     String("Office Supplies"),
		Unused:          Float64(0),
		Status:          String("created"),
		Currency:        String("USD"),
		AgentId:         Int64(11),
		LocationId:      Int64(61),
		LedgerAccountId: Int64(4790),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("CashReceipt.Update returned %+v, want %+v", cc, want)
	}
}

func TestCashReceiptService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cashReceipts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.CashReceipt.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("CashReceipt.Delete returned error: %v", err)
	}
}
