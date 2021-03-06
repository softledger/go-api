package softledger

import (
	"context"
	"fmt"
)

type WalletService service

type Wallet struct {
	ID                  *string        `json:"_id,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Description         *string        `json:"description,omitempty"`
	Address             *string        `json:"address,omitempty"`
	AddressTag          *string        `json:addressTag,omitempty"`
	Exchange            *string        `json:"exchange,omitempty"`
	DepositAccountId    *int64         `json:"DepositAccountId,omitempty"`
	WithdrawalAccountId *int64         `json:"WithdrawalAccountId,omitempty"`
	DepositAccount      *LedgerAccount `json:"DepositAccount,omitempty"`
	WithdrawalAccount   *LedgerAccount `json:"WithdrawalAccount,omitempty"`
}

type walletResponse struct {
	Data       []*Wallet `json:"data"`
	TotalItems int       `json:"totalItems"`
}

func (c Wallet) String() string {
	return Stringify(c)
}

func (s *WalletService) All(ctx context.Context, qry *QueryParams) ([]*Wallet, int, *Response, error) {

	u, err := addParams("/crypto/wallets", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var wallets *walletResponse
	resp, err := s.client.Do(ctx, req, &wallets)
	if err != nil {
		return nil, 0, resp, err
	}

	return wallets.Data, wallets.TotalItems, resp, nil

}

func (s *WalletService) One(ctx context.Context, _id *string) (*Wallet, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/wallets", *_id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Wallet
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *WalletService) Create(ctx context.Context, payload *Wallet) (*Wallet, *Response, error) {

	u := fmt.Sprintf("/crypto/wallets")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Wallet
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *WalletService) Update(ctx context.Context, _id int64, payload *Wallet) (*Wallet, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/wallets", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Wallet
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *WalletService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/crypto/wallets", _id)

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
