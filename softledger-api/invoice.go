package softledger

import (
	"context"
	"fmt"
)

type InvoiceService service

type Invoice struct {
	ID                *int64   `json:"_id"`
	Description       *string  `json:"description"`
	DueDate           *string  `json:"dueDate"`
	PostedDate        *string  `json:"postedDate"`
	InvoiceDate       *string  `json:"invoiceDate"`
	Amount            *float64 `json:"amount"`
	AmountPayable     *float64 `json:"amountPayable"`
	Status            *string  `json:"status"`
	Currency          *string  `json:"currency"`
	Notes             *string  `json:"notes"`
	Number            *string  `json:"number"`
	TemplateId        *int64
	LocationId        *int64
	Location          *Location
	ICLocationId      *int64
	ICLocation        *Location
	AgentId           *int64
	Agent             *Customer
	ARAccountId       *int64
	ARAccount         *LedgerAccount
	ShippingAddressId *int64
	ShippingAddress   *Address
	BillingAddressId  *int64
	BillingAddress    *Address
	Attachments       *Attachments
	InvoiceLineItems  *InvoiceLineItems
}

type InvoiceLineItem struct {
	Description     *string  `json:"description"`
	Quantity        *float64 `json:"quantity"`
	CostCenterId    *int64
	CostCenter      *CostCenter
	KitId           *int64
	Kit             *Kit
	JobId           *int64
	Job             *Job
	UnitAmount      *float64 `json:"unitAmount"`
	ItemId          *int64
	Item            *Item
	LedgerAccountId *int64
	LedgerAccount   *LedgerAccount
	KitLineItems    []*InvoiceLineItem
}

type InvoiceLineItems []*InvoiceLineItem

type invoiceResponse struct {
	Data       []*Invoice `json:"data"`
	TotalItems int        `json:"totalItems"`
}

func (c Invoice) String() string {
	return Stringify(c)
}

func (s *InvoiceService) All(ctx context.Context, qry *QueryParams) ([]*Invoice, int, *Response, error) {

	u, err := addParams("invoices", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var invoices *invoiceResponse
	resp, err := s.client.Do(ctx, req, &invoices)
	if err != nil {
		return nil, 0, resp, err
	}

	return invoices.Data, invoices.TotalItems, resp, nil

}

func (s *InvoiceService) One(ctx context.Context, _id int64) (*Invoice, *Response, error) {

	u := fmt.Sprintf("%v/%v", "invoices", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Invoice
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *InvoiceService) Create(ctx context.Context, payload *Invoice) (*Invoice, *Response, error) {

	u := fmt.Sprintf("invoices")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Invoice
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *InvoiceService) Update(ctx context.Context, _id int64, payload *Invoice) (*Invoice, *Response, error) {

	u := fmt.Sprintf("%v/%v", "invoices", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Invoice
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *InvoiceService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "invoices", _id)

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

func (s *InvoiceService) Issue(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v/issue", "invoices", _id)

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

func (s *InvoiceService) Void(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v/void", "invoices", _id)

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

func (s *InvoiceService) Apply(ctx context.Context, payload *interface{}) (*Response, error) {

	u := fmt.Sprintf("invoices/apply")

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
