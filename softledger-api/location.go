package softledger

import (
	"context"
	"fmt"
)

type LocationService service

type Location struct {
	ID          *int64      `json:"_id"`
	Id          *string     `json:"id"`
	Name        *string     `json:"name"`
	Currency    *string     `json:"currency"`
	Description *string     `json:"description"`
	Parent_id   *int64      `json:"parent_id"`
	Children    []*Location `json:"children"`
}

type locationResponse struct {
	Data       []*Location `json:"data"`
	TotalItems int         `json:"totalItems"`
}

func (cc Location) String() string {
	return Stringify(cc)
}

func (s *LocationService) All(ctx context.Context, qry *QueryParams) ([]*Location, int, *Response, error) {

	u, err := addParams("/locations", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *locationResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *LocationService) One(ctx context.Context, _id int64) (*Location, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/locations", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Location
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *LocationService) Create(ctx context.Context, payload *Location) (*Location, *Response, error) {

	u := fmt.Sprintf("/locations")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Location
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *LocationService) Update(ctx context.Context, _id int64, payload *Location) (*Location, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/locations", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Location
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *LocationService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/locations", _id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
