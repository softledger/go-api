package softledger

import (
	"context"
	"fmt"
)

type CustomerService service

type Customer struct {
	ID          *int64  `json:"_id"`
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Description *string `json:"description"`
	Website     *string `json:"website"`
	Addresses   *Addresses
	Contacts    *Contacts
}

type customerResponse struct {
	Data       []*Customer `json:"data"`
	TotalItems int         `json:"totalItems"`
}

func (c Customer) String() string {
	return Stringify(c)
}

func (s *CustomerService) All(ctx context.Context, qry *QueryParams) ([]*Customer, int, *Response, error) {

	u, err := addParams("customers", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var customers *customerResponse
	resp, err := s.client.Do(ctx, req, &customers)
	if err != nil {
		return nil, 0, resp, err
	}

	return customers.Data, customers.TotalItems, resp, nil

}

func (s *CustomerService) One(ctx context.Context, _id int64) (*Customer, *Response, error) {

	u := fmt.Sprintf("%v/%v", "customers", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Customer
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *CustomerService) Create(ctx context.Context, payload *Customer) (*Customer, *Response, error) {

	u := fmt.Sprintf("customers")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Customer
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CustomerService) Update(ctx context.Context, _id int64, payload *Customer) (*Customer, *Response, error) {

	u := fmt.Sprintf("%v/%v", "customers", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Customer
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CustomerService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "customers", _id)

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
