package softledger

import (
	"context"
	"fmt"
)

type SettingsService service

type Settings struct {
	Timezone                   *string `json:"timezone"`
	DateFormat                 *string `json:"dateFormat"`
	CcInvoice                  *bool   `json:"ccInvoice"`
	CostingMethod              *bool   `json:"costingMethod"`
	EmailOnInvoiceIssued       *bool   `json:"emailOnInvoiceIssued"`
	EmailOnSalesQuoteIssued    *bool   `json:"emailOnSalesQuoteIssued"`
	EmailOnPurchaseOrderIssued *bool   `json:"emailOnPurchaseOrderIssued"`
	UseLocationOnDocuments     *bool   `json:"useLocationOnDocuments"`
	PostCRJournal              *bool   `json:"postCRJournal"`
	DisplayItem                *string `json:"displayItem"`
	ConfirmDelete              *bool   `json:"confirmDelete"`

	//default accounts
	DefaultAccountsReceivable *int64 `json:"defaultAccountsReceivable"`
	DefaultRev                *int64 `json:"defaultRev"`
	DefaultAccountsPayable    *int64 `json:"defaultAccountsPayable"`
	DefaultCash               *int64 `json:"defaultCash"`
	DefaultInventoryAccrual   *int64 `json:"defaultInventoryAccrual"`
	DefaultItemInventoryAsset *int64 `json:"defaultItemInventoryAsset"`
	DefaultItemCOGS           *int64 `json:"defaultItemCOGS"`
	DefaultUnappliedCash      *int64 `json:"defaultUnappliedCash"`
}

func (s *SettingsService) Get(ctx context.Context) (*Settings, *Response, error) {

	u, err := addParams("/settings", nil)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var ccs *Settings
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, resp, err
	}

	return ccs, resp, nil

}

func (s *SettingsService) Update(ctx context.Context, payload *Settings) (*Settings, *Response, error) {

	u := fmt.Sprintf("/settings")

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Settings
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}
