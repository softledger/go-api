package softledger

import (
	"context"
	"fmt"
)

type BillService service

type Bill struct {
	ID              *int64   `json:"_id"`
	Description     *string  `json:"description"`
	DueDate         *string  `json:"dueDate"`
	PostingDate     *string  `json:"postingDate"`
	InvoiceDate     *string  `json:"invoiceDate"`
	Amount          *float64 `json:"amount"`
	DueAmount       *float64 `json:"dueAmount"`
	Status          *string  `json:"status"`
	Currency        *string  `json:"currency"`
	Notes           *string  `json:"notes"`
	InvoiceNumber   *string  `json:"invoiceNumber"`
	LocationId      *int64
	Location        *Location
	ICLocationId    *int64
	ICLocation      *Location
	VendorId        *int64
	Vendor          *Vendor
	APAccountId     *int64
	APAccount       *LedgerAccount
	PurchaseOrderId *int64
	PurchaseOrder   *PurchaseOrder
	Attachments     *Attachments
	BillLineItems   *BillLineItems `json:"billLineItems"`
}

type BillLineItem struct {
	Description     *string  `json:"description"`
	Quantity        *float64 `json:"quantity"`
	CostCenterId    *int64
	CostCenter      *CostCenter
	KitId           *int64
	Kit             *Kit
	JobId           *int64
	Job             *Job
	Amount          *float64 `json:"amount"`
	ItemId          *int64
	Item            *Item
	LedgerAccountId *int64
	LedgerAccount   *LedgerAccount
	KitLineItems    []*BillLineItem
}

type BillLineItems []*BillLineItem

type billResponse struct {
	Data       []*Bill `json:"data"`
	TotalItems int     `json:"totalItems"`
}

func (c Bill) String() string {
	return Stringify(c)
}

func (s *BillService) All(ctx context.Context, qry *QueryParams) ([]*Bill, int, *Response, error) {

	u, err := addParams("bills", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var bills *billResponse
	resp, err := s.client.Do(ctx, req, &bills)
	if err != nil {
		return nil, 0, resp, err
	}

	return bills.Data, bills.TotalItems, resp, nil

}

func (s *BillService) One(ctx context.Context, _id int64) (*Bill, *Response, error) {

	u := fmt.Sprintf("%v/%v", "bills", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Bill
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *BillService) Create(ctx context.Context, payload *Bill) (*Bill, *Response, error) {

	u := fmt.Sprintf("bills")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Bill
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *BillService) Update(ctx context.Context, _id int64, payload *Bill) (*Bill, *Response, error) {

	u := fmt.Sprintf("%v/%v", "bills", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Bill
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *BillService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "bills", _id)

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

func (s *BillService) Approve(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "bills/approve", _id)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (s *BillService) Void(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "bills/void", _id)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

func (s *BillService) Pay(ctx context.Context, payload *interface{}) (*Response, error) {

	u := fmt.Sprintf("bills/pay")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
