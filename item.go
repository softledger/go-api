package softledger

import (
	"context"
	"fmt"
)

type ItemService service

type Item struct {
	ID                   *int64   `json:"_id"`
	Number               *string  `json:"number"`
	Name                 *string  `json:"name"`
	Sku                  *string  `json:"sku"`
	Description          *string  `json:"description"`
	SalesPrice           *float64 `json:"salesPrice"`
	InvoiceAccountId     *int64   `json:"InvoiceAccountId"`
	InvoiceAccount       *LedgerAccount
	BillAccountId        *int64 `json:"BillAccountId"`
	BillAccount          *LedgerAccount
	InventoryAccountId   *int64 `json:"InventoryAccountId"`
	InventoryAccount     *LedgerAccount
	CogsAccountId        *int64 `json:"CogsAccountId"`
	CogsAccount          *LedgerAccount
	LowStockNotification *bool   `json:"lowStockNotification"`
	LowStockThreshold    *int64  `json:"lowStockThreshold"`
	LowStockEmail        *string `json:"lowStockEmail"`
	//CustomFields         *interface{} `json:"customFields"`
}

type itemResponse struct {
	Data       []*Item `json:"data"`
	TotalItems int     `json:"totalItems"`
}

func (cc Item) String() string {
	return Stringify(cc)
}

func (s *ItemService) All(ctx context.Context, qry *QueryParams) ([]*Item, int, *Response, error) {

	u, err := addParams("/items", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *itemResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *ItemService) One(ctx context.Context, _id int64) (*Item, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/items", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Item
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *ItemService) Create(ctx context.Context, payload *Item) (*Item, *Response, error) {

	u := fmt.Sprintf("/items")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Item
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *ItemService) Update(ctx context.Context, _id int64, payload *Item) (*Item, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/items", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Item
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *ItemService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/items", _id)

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
