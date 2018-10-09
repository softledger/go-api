package softledger

import (
	"context"
	"fmt"
)

type LedgerAccountService service

type LedgerAccount struct {
	ID                      *int64  `json:"_id"`
	Name                    *string `json:"name"`
	NaturalBalance          *string `json:"naturalBalance"`
	Description             *string `json:"description"`
	Number                  *string `json:"number"`
	Type                    *string `json:"type"`
	Subtype                 *string `json:"subtype"`
	IncludeLocationChildren *bool   `json:"includeLocationChildren"`
	LocationId              *int64
	Location                *Location
}

type ledgerAccountResponse struct {
	Data       []*LedgerAccount `json:"data"`
	TotalItems int              `json:"totalItems"`
}

func (cc LedgerAccount) String() string {
	return Stringify(cc)
}

func (s *LedgerAccountService) All(ctx context.Context, qry *QueryParams) ([]*LedgerAccount, int, *Response, error) {

	u, err := addParams("/ledger_accounts", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *ledgerAccountResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *LedgerAccountService) One(ctx context.Context, _id int64) (*LedgerAccount, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/ledger_accounts", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *LedgerAccount
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *LedgerAccountService) Create(ctx context.Context, payload *LedgerAccount) (*LedgerAccount, *Response, error) {

	u := fmt.Sprintf("/ledger_accounts")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *LedgerAccount
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *LedgerAccountService) Update(ctx context.Context, _id int64, payload *LedgerAccount) (*LedgerAccount, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/ledger_accounts", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *LedgerAccount
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *LedgerAccountService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/ledger_accounts", _id)

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
