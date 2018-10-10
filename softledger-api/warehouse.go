package softledger

import (
	"context"
	"fmt"
)

type WarehouseService service

type Warehouse struct {
	ID          *int64  `json:"_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Addresses   *Addresses
}

type warehouseResponse struct {
	Data       []*Warehouse `json:"data"`
	TotalItems int          `json:"totalItems"`
}

func (c Warehouse) String() string {
	return Stringify(c)
}

func (s *WarehouseService) All(ctx context.Context, qry *QueryParams) ([]*Warehouse, int, *Response, error) {

	u, err := addParams("warehouses", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var warehouses *warehouseResponse
	resp, err := s.client.Do(ctx, req, &warehouses)
	if err != nil {
		return nil, 0, resp, err
	}

	return warehouses.Data, warehouses.TotalItems, resp, nil

}

func (s *WarehouseService) One(ctx context.Context, _id int64) (*Warehouse, *Response, error) {

	u := fmt.Sprintf("%v/%v", "warehouses", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Warehouse
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *WarehouseService) Create(ctx context.Context, payload *Warehouse) (*Warehouse, *Response, error) {

	u := fmt.Sprintf("warehouses")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Warehouse
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *WarehouseService) Update(ctx context.Context, _id int64, payload *Warehouse) (*Warehouse, *Response, error) {

	u := fmt.Sprintf("%v/%v", "warehouses", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Warehouse
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *WarehouseService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "warehouses", _id)

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
