package softledger

import (
	"context"
	"fmt"
)

type CoinService service

type Coin struct {
	ID                *string        `json:"_id,omitempty"`
	Name              *string        `json:"name,omitempty"`
	Symbol            *string        `json:"symbol,omitempty"`
	AssetAccount      *LedgerAccount `json:"AssetAccount,omitempty"`
	FeeAccount        *LedgerAccount `json:"FeeAccount,omitempty"`
	LTGainLossAccount *LedgerAccount `json:"LTGainLossAccount,omitempty"`
	STGainLossAccount *LedgerAccount `json:"STGainLossAccount,omitempty"`
}

type coinResponse struct {
	Data       []*Coin `json:"data"`
	TotalItems int     `json:"totalItems"`
}

func (c Coin) String() string {
	return Stringify(c)
}

func (s *CoinService) All(ctx context.Context, qry *QueryParams) ([]*Coin, int, *Response, error) {

	u, err := addParams("/crypto/coins", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var coins *coinResponse
	resp, err := s.client.Do(ctx, req, &coins)
	if err != nil {
		return nil, 0, resp, err
	}

	return coins.Data, coins.TotalItems, resp, nil
}

func (s *CoinService) One(ctx context.Context, _id int64) (*Coin, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/coins", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Coin
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *CoinService) Create(ctx context.Context, payload *Coin) (*Coin, *Response, error) {

	u := fmt.Sprintf("/crypto/coins")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Coin
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *CoinService) Hide(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/coins", _id)

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
