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

func TestJournalService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
			"_id":1,
			"entryType":"1",
			"number":1,
			"sourceLedger": "1"
		}], "totalItems": 1}`)
	})

	ccs, totalItems, _, err := client.Journal.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Journal.All returned error: %v", err)
	}

	if totalItems != 1 {
		t.Errorf("Wrong sourceLedger of journals, want %v, got %v", 1, totalItems)
	}

	want := []*Journal{{
		ID:           Int64(1),
		EntryType:    String("1"),
		Number:       Int64(1),
		SourceLedger: String("1"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Journal.All returned %+v, want %+v", ccs, want)
	}

}

func TestJournalService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"_id":1,
			"entryType":"1",
			"number": 1,
			"sourceLedger": "1"
		}`)
	})

	cc, _, err := client.Journal.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Journal.One returned error: %v", err)
	}

	want := &Journal{
		ID:           Int64(1),
		EntryType:    String("1"),
		Number:       Int64(1),
		SourceLedger: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Journal.One returned %+v, want %+v", cc, want)
	}

}

func TestJournalService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
			"_id":1,
			"entryType":"1",
			"number": 1,
			"sourceLedger": "1"
		}`)
	})

	payload := &Journal{
		EntryType: String("1"),
		Number:    Int64(1),
	}

	cc, _, err := client.Journal.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Journal.Create returned error: %v", err)
	}

	want := &Journal{
		ID:           Int64(1),
		EntryType:    String("1"),
		Number:       Int64(1),
		SourceLedger: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Journal.Create returned %+v, want %+v", cc, want)
	}
}

func TestJournalService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"_id":1,
			"entryType":"1",
			"number": 1,
			"sourceLedger": "1"
		}`)
	})

	payload := &Journal{
		EntryType: String("1"),
	}

	cc, _, err := client.Journal.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Journal.Update returned error: %v", err)
	}

	want := &Journal{
		ID:           Int64(1),
		EntryType:    String("1"),
		Number:       Int64(1),
		SourceLedger: String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Journal.Update returned %+v, want %+v", cc, want)
	}
}

func TestJournalService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Journal.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Journal.Delete returned error: %v", err)
	}
}

func TestJournalService_post(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/journals/1/post", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Journal.Post(context.Background(), 1)
	if err != nil {
		t.Errorf("Journal.Post returned error: %v", err)
	}
}
