package softledger

import (
	"context"
	"fmt"
	"time"
)

type CryptoTransactionService service

type CryptoTransaction struct {
	ID             *int64     `json:"_id,omitempty"`
	Date           *time.Time `json:"date,omitempty"`
	Type           *string    `json:"type,omitempty"`
	Locked         *bool      `json:"locked,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
	Currency       *string    `json:"currency,omitempty"`
	ExternalId     *string    `json:"externalId,omitempty"`
	ExternalSource *string    `json:"externalSource,omitempty"`
	TxHash         *string    `json:"txHash,omitempty"`

	RQty *string `json:"rQty,omitempty"`
	SQty *string `json:"sQty,omitempty"`
	FQty *string `json:"fQty,omitempty"`

	RCoinId *string `json:"rCoinId,omitempty"`
	SCoinId *string `json:"sCoinId,omitempty"`
	FCoinId *string `json:"fCoinId,omitempty"`

	FCoin *Coin `json:"fCoin,omitempty"`
	SCoin *Coin `json:"sCoin,omitempty"`
	RCoin *Coin `json:"rCoin,omitempty"`

	RWalletId *string `json:"rWalletId,omitempty"`
	SWalletId *string `json:"sWalletId,omitempty"`
	FWalletId *string `json:"fWalletId,omitempty"`

	RPrice *string `json:"rPrice,omitempty"`
	SPrice *string `json:"sPrice,omitempty"`
	FPrice *string `json:"fPrice,omitempty"`

	SCostBasis *string `json:"sCostBasis,omitempty"`
	FCostBasis *string `json:"fCostBasis,omitempty"`

	SCostLayers []*CryptoTransactionCostLayer `json:"sCostLayers,omitempty"`
	FCostLayers []*CryptoTransactionCostLayer `json:"fCostLayers,omitempty"`

	LedgerAccountId *int64  `json:"LedgerAccountId,omitempty"`
	CustomerId      *int64  `json:"CustomerId,omitempty"`
	JournalId       *string `json:"JournalId,omitempty"`

	//not settable in create/update
	Error        *CryptoTransactionError `json:"error,omitempty"`
	QtyPicked    *string                 `json:"qtyPicked,omitempty"`
	CurrencyRate *string                 `json:"currencyRate,omitempty"`
}

type CryptoTransactionError struct {
	Type *string `json:"type,omitempty"`
	Msg  *string `json:"msg,omitempty"`
}

type CryptoTransactionCostLayer struct {
	ID        *int64     `json:"_id,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
	CostBasis *string    `json:"costBasis,omitempty"`
	QtyPicked *string    `json:"qtyPicked,omitempty"`
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

	u := fmt.Sprintf("/crypto")

	req, err := s.client.NewSvcRequest("POST", u, payload)
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
