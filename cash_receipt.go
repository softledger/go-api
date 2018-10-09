package softledger

import (
	"context"
	"fmt"
)

type CashReceiptService service

type CashReceipt struct {
	ID                     *int64   `json:"_id"`
	Number                 *string  `json:"number"`
	Type                   *string  `json:"type"`
	Amount                 *float64 `json:"amount"`
	Unused                 *float64 `json:"unused"`
	Description            *string  `json:"description"`
	Attachments            *Attachments
	ReceiveDate            *string `json:"receiveDate"`
	PostingDate            *string `json:"postingDate"`
	Status                 *string `json:"status"`
	LedgerAccountId        *int64
	LedgerAccount          *LedgerAccount
	UnappliedCashAccountId *int64
	UnappliedCashAccount   *LedgerAccount
	Currency               *string `json:"currency"`
	LocationId             *int64
	Location               *Location
	AgentId                *int64
	Agent                  *Customer
}

type cashReceiptResponse struct {
	Data       []*CashReceipt `json:"data"`
	TotalItems int            `json:"totalItems"`
}

func (c CashReceipt) String() string {
	return Stringify(c)
}

func (s *CashReceiptService) All(ctx context.Context, qry *QueryParams) ([]*CashReceipt, int, *Response, error) {

	u, err := addParams("cashReceipts", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var cashReceipts *cashReceiptResponse
	resp, err := s.client.Do(ctx, req, &cashReceipts)
	if err != nil {
		return nil, 0, resp, err
	}

	return cashReceipts.Data, cashReceipts.TotalItems, resp, nil

}

func (s *CashReceiptService) One(ctx context.Context, _id int64) (*CashReceipt, *Response, error) {

	u := fmt.Sprintf("%v/%v", "cashReceipts", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *CashReceipt
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *CashReceiptService) Create(ctx context.Context, payload *CashReceipt) (*CashReceipt, *Response, error) {

	u := fmt.Sprintf("cashReceipts")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CashReceipt
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CashReceiptService) Update(ctx context.Context, _id int64, payload *CashReceipt) (*CashReceipt, *Response, error) {

	u := fmt.Sprintf("%v/%v", "cashReceipts", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CashReceipt
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CashReceiptService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "cashReceipts", _id)

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
