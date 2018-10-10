package softledger

import (
	"context"
	"fmt"
)

type VendorService service

type Vendor struct {
	ID          *int64  `json:"_id"`
	Name        *string `json:"name"`
	ShortName   *string `json:"shortName"`
	NameOnCheck *string `json:"nameOnCheck"`
	AccNumber   *string `json:"accNumber"`
	Email       *string `json:"email"`
	Is1099      *bool   `json:"is1099"`
	EIN         *string
	Addresses   *Addresses
	Contacts    *Contacts
}

type vendorResponse struct {
	Data       []*Vendor `json:"data"`
	TotalItems int       `json:"totalItems"`
}

func (c Vendor) String() string {
	return Stringify(c)
}

func (s *VendorService) All(ctx context.Context, qry *QueryParams) ([]*Vendor, int, *Response, error) {

	u, err := addParams("vendors", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var vendors *vendorResponse
	resp, err := s.client.Do(ctx, req, &vendors)
	if err != nil {
		return nil, 0, resp, err
	}

	return vendors.Data, vendors.TotalItems, resp, nil

}

func (s *VendorService) One(ctx context.Context, _id int64) (*Vendor, *Response, error) {

	u := fmt.Sprintf("%v/%v", "vendors", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Vendor
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *VendorService) Create(ctx context.Context, payload *Vendor) (*Vendor, *Response, error) {

	u := fmt.Sprintf("vendors")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Vendor
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *VendorService) Update(ctx context.Context, _id int64, payload *Vendor) (*Vendor, *Response, error) {

	u := fmt.Sprintf("%v/%v", "vendors", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Vendor
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *VendorService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "vendors", _id)

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
