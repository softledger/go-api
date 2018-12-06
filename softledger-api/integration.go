package softledger

import (
	"context"
	"fmt"
)

type IntegrationService service

type Integration struct {
	//uniqueness is ID only, as we can have
	ID       *string `json:"_id,omitempty"`
	TenantId *string `json"tenantId,omitempty"`
	Type     *string `json:"type,omitempty"`
	Name     *string `json:"name,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty"`
	Details  *Detail `json:"details,omitempty"`
}

type Detail struct {
	LedgerAccountId *int64  `json:"LedgerAccountId,omitempty"`
	WalletId        *string `json:"WalletId,omitempty"`
	ApiKey          *string `json:"ApiKey,omitempty"`
	ApiSecret       *string `json:"ApiSecret,omitempty"`
}

type integrationResponse struct {
	Data       []*Integration `json:"data"`
	TotalItems int            `json:"totalItems"`
}

func (c Integration) String() string {
	return Stringify(c)
}

func (s *IntegrationService) All(ctx context.Context) ([]*Integration, *Response, error) {

	u := fmt.Sprintf("%v", "integrations")

	req, err := s.client.NewSvcRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var integrations []*Integration
	resp, err := s.client.Do(ctx, req, &integrations)
	if err != nil {
		return nil, resp, err
	}

	return integrations, resp, nil

}

func (s *IntegrationService) One(ctx context.Context, _id *string) (*Integration, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/integrations", *_id)

	req, err := s.client.NewSvcRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Integration
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *IntegrationService) Create(ctx context.Context, payload *Integration) (*Integration, *Response, error) {

	u := fmt.Sprintf("integrations")

	req, err := s.client.NewSvcRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Integration
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *IntegrationService) Update(ctx context.Context, _id int64, payload *Integration) (*Integration, *Response, error) {

	u := fmt.Sprintf("%v/%v", "integrations", _id)

	req, err := s.client.NewSvcRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Integration
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *IntegrationService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "integrations", _id)

	req, err := s.client.NewSvcRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
