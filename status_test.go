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

func TestStatusService_read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/status/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		fmt.Fprint(w, `{
			"type": "costbasis",
			"status": "running",
			"data": {
				"_id": "1234",
				"progress": "53.2",
				"timestamp": "10/14/2017",
				"err": {
					"type": "CostBasisError",
					"msg": "big bad error"
				}
			}
		}`)
	})

	cc, _, err := client.Status.Read(context.Background(), "1")
	if err != nil {
		t.Errorf("Status.Read returned error: %v", err)
	}

	want := &Status{
		Type:   String("costbasis"),
		Status: String("running"),
		Data: &StatusData{
			ID:        String("1234"),
			Progress:  String("53.2"),
			Timestamp: String("10/14/2017"),
			Err: &StatusErr{
				Type: String("CostBasisError"),
				Msg:  String("big bad error"),
			},
		},
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Status.Read returned %+v, want %+v", cc, want)
	}

}

func TestStatusService_write(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/status/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		fmt.Fprint(w, `{
			"type": "costbasis",
			"status": "running",
			"data": {
				"_id": "1234",
				"progress": "53.2",
				"timestamp": "10/14/2017",
				"err": {
					"type": "CostBasisError",
					"msg": "big bad error"
				}
			}
		}`)
	})

	payload := &Status{
		Status: String("running"),
		Data: &StatusData{
			ID:        String("1234"),
			Progress:  String("53.2"),
			Timestamp: String("10/14/2017"),
			Err: &StatusErr{
				Type: String("CostBasisError"),
				Msg:  String("big bad error"),
			},
		},
	}

	cc, _, err := client.Status.Write(context.Background(), "1", payload)
	if err != nil {
		t.Errorf("Status.Write returned error: %v", err)
	}

	want := &Status{
		Type:   String("costbasis"),
		Status: String("running"),
		Data: &StatusData{
			ID:        String("1234"),
			Progress:  String("53.2"),
			Timestamp: String("10/14/2017"),
			Err: &StatusErr{
				Type: String("CostBasisError"),
				Msg:  String("big bad error"),
			},
		},
	}
	if !reflect.DeepEqual(cc, want) {
		t.Errorf("Status.Write returned %+v, want %+v", cc, want)
	}
}
