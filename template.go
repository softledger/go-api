package softledger

import (
	"context"
	"fmt"
)

type TemplateService service

type Template struct {
	ID   *int64       `json:"_id"`
	Name *string      `json:"name"`
	Type *string      `json:"type"`
	Data *interface{} `json:"data"`
}

type templateResponse struct {
	Data       []*Template `json:"data"`
	TotalItems int         `json:"totalItems"`
}

func (c Template) String() string {
	return Stringify(c)
}

func (s *TemplateService) All(ctx context.Context, qry *QueryParams) ([]*Template, int, *Response, error) {

	u, err := addParams("templates", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var templates *templateResponse
	resp, err := s.client.Do(ctx, req, &templates)
	if err != nil {
		return nil, 0, resp, err
	}

	return templates.Data, templates.TotalItems, resp, nil

}

func (s *TemplateService) One(ctx context.Context, _id int64) (*Template, *Response, error) {

	u := fmt.Sprintf("%v/%v", "templates", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Template
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *TemplateService) Create(ctx context.Context, payload *Template) (*Template, *Response, error) {

	u := fmt.Sprintf("templates")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Template
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *TemplateService) Update(ctx context.Context, _id int64, payload *Template) (*Template, *Response, error) {

	u := fmt.Sprintf("%v/%v", "templates", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Template
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *TemplateService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "templates", _id)

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
