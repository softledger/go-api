package softledger

import (
	"context"
	"fmt"
)

type JournalService service

type Journal struct {
	ID           *int64  `json:"_id"`
	Number       *string `json:"number"`
	Status       *string `json:"status"`
	EntryType    *string `json:"entryType"`
	SourceLedger *string `json:"sourceLedger"`
	Reference    *string `json:"reference"`
	Transactions *Transactions
}

type Transactions []*Transaction

type Transaction struct {
	TransactionDate *string  `json:"transactionDate"`
	PostedDate      *string  `json:"postedDate"`
	Debit           *float64 `json:"debit"`
	Credit          *float64 `json:"credit"`

	CostCenterId    *int64
	CostCenter      *CostCenter
	LedgerAccountId *int64
	LedgerAccount   *LedgerAccount
	JobId           *int64
	Job             *Job
	LocationId      *int64
	Location        *Location
	ICLocationId    *int64
	ICLocation      *Location
	InvoiceId       *int64
	Invoice         *Invoice
	BillId          *int64
	Bill            *Bill
	AgentId         *int64
	Agent           *Customer
	VendorId        *int64
	Vendor          *Vendor
	CashReceiptId   *int64
	CashReceipt     *CashReceipt
}

type journalResponse struct {
	Data       []*Journal `json:"data"`
	TotalItems int        `json:"totalItems"`
}

func (cc Journal) String() string {
	return Stringify(cc)
}

func (s *JournalService) All(ctx context.Context, qry *QueryParams) ([]*Journal, int, *Response, error) {

	u, err := addParams("/journals", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *journalResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalItems, resp, nil

}

func (s *JournalService) One(ctx context.Context, _id int64) (*Journal, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/journals", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Journal
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *JournalService) Create(ctx context.Context, payload *Journal) (*Journal, *Response, error) {

	u := fmt.Sprintf("/journals")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Journal
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *JournalService) Update(ctx context.Context, _id int64, payload *Journal) (*Journal, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/journals", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Journal
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *JournalService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/journals", _id)

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

func (s *JournalService) Post(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v/post", "/journals", _id)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
