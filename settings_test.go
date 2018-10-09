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

func TestSettingsService_get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"timezone":"US/Pacific"}`)
	})

	cc, _, err := client.Settings.Get(context.Background())
	if err != nil {
		t.Errorf("Settings.One returned error: %v", err)
	}

	want := &Settings{
		Timezone: String("US/Pacific"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Settings.One returned %+v, want %+v", cc, want)
	}

}

func TestSettingsService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{"timezone":"US/Pacific"}`)
	})

	payload := &Settings{
		Timezone: String("US/Pacific"),
	}

	cc, _, err := client.Settings.Update(context.Background(), payload)
	if err != nil {
		t.Errorf("Settings.Update returned error: %v", err)
	}

	want := &Settings{
		Timezone: String("US/Pacific"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Settings.Update returned %+v, want %+v", cc, want)
	}
}
