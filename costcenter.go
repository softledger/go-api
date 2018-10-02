package softledger

import (
	"context"
	"fmt"
)

const (
	BASE = "/cost_centers"
)

type CostCenterService service

type CostCenter struct {
	ID          *int64  `json:"_id"`
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type ccResponse struct {
	Data       []*CostCenter `json:"data"`
	TotalItems int           `json:"totalItems"`
}

func (cc CostCenter) String() string {
	return Stringify(cc)
}

func (s *CostCenterService) All(ctx context.Context, qry *QueryParams) ([]*CostCenter, int, *Response, error) {

	u, err := addParams(BASE, qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *ccResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *CostCenterService) One(ctx context.Context, _id int64) (*CostCenter, *Response, error) {

	u := fmt.Sprintf("%v/%v", BASE, _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *CostCenter
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *CostCenterService) Create(ctx context.Context, payload *CostCenter) (*CostCenter, *Response, error) {

	u := fmt.Sprintf(BASE)

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CostCenter
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CostCenterService) Update(ctx context.Context, _id int64, payload *CostCenter) (*CostCenter, *Response, error) {

	u := fmt.Sprintf("%v/%v", BASE, _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CostCenter
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CostCenterService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", BASE, _id)

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
