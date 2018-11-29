package softledger

import (
	"context"
	"fmt"
)

type CurrencyService service

type Currency struct {
	Fraction       *int64  `json:"fraction"`
	RoundingMethod *int64  `json:"rounding_method"`
	Code           *string `json:"code"`
	Name           *string `json:"name"`
	Symbol         *string `json:"symbol"`
}

func (cc Currency) String() string {
	return Stringify(cc)
}

func (s *CurrencyService) All(ctx context.Context) ([]*Currency, *Response, error) {

	u := fmt.Sprintf("%v", "/currency")

	req, err := s.client.NewSvcRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc []*Currency
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}
