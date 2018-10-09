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

func TestJobService_all(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{"data":[{
			"_id":1,
			"description":"1",
			"name":"one",
			"number": "1"
		}], "totalJobs": 1}`)
	})

	ccs, totalJobs, _, err := client.Job.All(context.Background(), nil)
	if err != nil {
		t.Errorf("Job.All returned error: %v", err)
	}

	if totalJobs != 1 {
		t.Errorf("Wrong number of jobs, want %v, got %v", 1, totalJobs)
	}

	want := []*Job{{
		ID:          Int64(1),
		Description: String("1"),
		Name:        String("one"),
		Number:      String("1"),
	}}
	if !reflect.DeepEqual(ccs, want) {
		t.Errorf("Job.All returned %+v, want %+v", ccs, want)
	}

}

func TestJobService_one(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"_id":1,
			"description":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	cc, _, err := client.Job.One(context.Background(), 1)
	if err != nil {
		t.Errorf("Job.One returned error: %v", err)
	}

	want := &Job{
		ID:          Int64(1),
		Description: String("1"),
		Name:        String("one"),
		Number:      String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Job.One returned %+v, want %+v", cc, want)
	}

}

func TestJobService_create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		fmt.Fprint(w, `{
			"_id":1,
			"description":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	payload := &Job{
		Description: String("1"),
		Name:        String("one"),
	}

	cc, _, err := client.Job.Create(context.Background(), payload)
	if err != nil {
		t.Errorf("Job.Create returned error: %v", err)
	}

	want := &Job{
		ID:          Int64(1),
		Description: String("1"),
		Name:        String("one"),
		Number:      String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Job.Create returned %+v, want %+v", cc, want)
	}
}

func TestJobService_update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"_id":1,
			"description":"1",
			"name":"one",
			"number": "1"
		}`)
	})

	payload := &Job{
		Description: String("1"),
	}

	cc, _, err := client.Job.Update(context.Background(), 1, payload)
	if err != nil {
		t.Errorf("Job.Update returned error: %v", err)
	}

	want := &Job{
		ID:          Int64(1),
		Description: String("1"),
		Name:        String("one"),
		Number:      String("1"),
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Job.Update returned %+v, want %+v", cc, want)
	}
}

func TestJobService_delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Job.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Job.Delete returned error: %v", err)
	}
}
