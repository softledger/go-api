package softledger

import (
	"context"
	"fmt"
)

type CryptoTransactionService service

type CryptoTransaction struct {
	ID     *string `json:"_id"`
	Date   *string `json:"date"`
	Type   *string `json:"type"`
	Locked *bool   `json:"locked"`
	Notes  *string `json:"notes"`

	RQty *float64 `json:"rQty"`
	SQty *float64 `json:"sQty"`
	FQty *float64 `json:"fQty"`

	RCoinId *int64 `json:"rCoinId"`
	SCoinId *int64 `json:"sCoinId"`
	FCoinId *int64 `json:"fCoinId"`

	RWalletId *int64 `json:"rWalletId"`
	SWalletId *int64 `json:"sWalletId"`
	FWalletId *int64 `json:"fWalletId"`

	RPrice *float64 `json:"rPrice"`
	SPrice *float64 `json:"sPrice"`
	FPrice *float64 `json:"fPrice"`

	SCostBasis *float64 `json:"sCostBasis"`
	FCostBasis *float64 `json:"fCostBasis"`

	SCostLayers *interface{} `json:"sCostLayers"`
	FCostLayers *interface{} `json:"fCostLayers"`

	LedgerAccount *LedgerAccount
	Customer      *Customer
}

type ctResponse struct {
	Data       []*CryptoTransaction `json:"data"`
	TotalItems int                  `json:"totalItems"`
}

func (cc CryptoTransaction) String() string {
	return Stringify(cc)
}

func (s *CryptoTransactionService) All(ctx context.Context, qry *QueryParams) ([]*CryptoTransaction, int, *Response, error) {

	u, err := addParams("/crypto/transactions", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *ctResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *CryptoTransactionService) One(ctx context.Context, _id int64) (*CryptoTransaction, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/transactions", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *CryptoTransaction
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *CryptoTransactionService) Create(ctx context.Context, payload *CryptoTransaction) (*CryptoTransaction, *Response, error) {

	u := fmt.Sprintf("/crypto/transactions")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CryptoTransaction
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CryptoTransactionService) Update(ctx context.Context, _id int64, payload *CryptoTransaction) (*CryptoTransaction, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/transactions", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *CryptoTransaction
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CryptoTransactionService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/transactions", _id)

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

func (s *CryptoTransactionService) Lock(ctx context.Context, payload *CryptoTransaction) (*Response, error) {

	u := fmt.Sprintf("%v/lock", "/crypto/transactions")

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, err
	}

	var cc *CryptoTransaction
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
