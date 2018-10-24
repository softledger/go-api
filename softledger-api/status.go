package softledger

import (
	"context"
	"fmt"
)

type StatusService service

type Status struct {
	Type   *string     `json:"type"`
	Status *string     `json:"status"`
	Data   *StatusData `json:"data"`
}

type StatusData struct {
	ID        *string    `json:"_id"`
	Err       *StatusErr `json:"err"`
	Progress  *string    `json:"progress"`
	Timestamp *string    `json:"timestamp"`
}

type StatusErr struct {
	Type *string `json:"type"`
	Msg  *string `json:"msg"`
}

func (cc Status) String() string {
	return Stringify(cc)
}

func (s *StatusService) Read(ctx context.Context, _type string) (*Status, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/status", _type)

	req, err := s.client.NewSvcRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Status
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *StatusService) Write(ctx context.Context, _type string, payload *Status) (*Status, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/status", _type)

	req, err := s.client.NewSvcRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Status
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}
