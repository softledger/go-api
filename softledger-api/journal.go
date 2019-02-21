package softledger

import (
	"context"
	"fmt"
	"time"
)

type JournalService service

type Journal struct {
	ID           *int64  `json:"_id,omitempty"`
	Number       *int64  `json:"number,omitempty"`
	Status       *string `json:"status,omitempty"`
	EntryType    *string `json:"entryType,omitempty"`
	SourceLedger *string `json:"sourceLedger,omitempty"`
	Reference    *string `json:"reference,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	Transactions []*Transaction
}

//type Transactions []*Transaction

type Transaction struct {
	TransactionDate *time.Time `json:"transactionDate,omitempty"`
	PostedDate      *time.Time `json:"postedDate,omitempty"`
	Debit           *string    `json:"debit,omitempty"`
	Credit          *string    `json:"credit,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Currency        *string    `json:"currency,omitempty"`

	CostCenterId    *int64         `json:",omitempty"`
	CostCenter      *CostCenter    `json:",omitempty"`
	LedgerAccountId *int64         `json:",omitempty"`
	LedgerAccount   *LedgerAccount `json:",omitempty"`
	JobId           *int64         `json:",omitempty"`
	Job             *Job           `json:",omitempty"`
	LocationId      *int64         `json:",omitempty"`
	Location        *Location      `json:",omitempty"`
	ICLocationId    *int64         `json:",omitempty"`
	ICLocation      *Location      `json:",omitempty"`
	InvoiceId       *int64         `json:",omitempty"`
	Invoice         *Invoice       `json:",omitempty"`
	BillId          *int64         `json:",omitempty"`
	Bill            *Bill          `json:",omitempty"`
	AgentId         *int64         `json:",omitempty"`
	Agent           *Customer      `json:",omitempty"`
	VendorId        *int64         `json:",omitempty"`
	Vendor          *Vendor        `json:",omitempty"`
	CashReceiptId   *int64         `json:",omitempty"`
	CashReceipt     *CashReceipt   `json:",omitempty"`
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
