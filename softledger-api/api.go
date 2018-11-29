package softledger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	//automatically adds api/ on the end
	defaultBaseURL = "https://api.softledger.com/"
	defaultSvcURL  = "https://services.softledger.com/"
	userAgent      = "go-softledger"
)

type Client struct {
	client *http.Client //HTTP client to talk to the api

	BaseURL *url.URL

	SvcURL *url.URL

	UserAgent string

	common service

	Bill              *BillService
	CashReceipt       *CashReceiptService
	Coin              *CoinService
	CostCenter        *CostCenterService
	CryptoTransaction *CryptoTransactionService
	Currency          *CurrencyService
	Customer          *CustomerService
	Invoice           *InvoiceService
	Item              *ItemService
	Job               *JobService
	Journal           *JournalService
	LedgerAccount     *LedgerAccountService
	Location          *LocationService
	Settings          *SettingsService
	Status            *StatusService
	Template          *TemplateService
	Vendor            *VendorService
	Warehouse         *WarehouseService
	Wallet            *WalletService
}

type QueryParams struct {
	Where  string `url:"where,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

type service struct {
	client *Client
}

func addParams(s string, params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		//nothing to add
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(params)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

func NewClient(httpClient *http.Client) *Client {
	//init default client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	svcURL, _ := url.Parse(defaultSvcURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		SvcURL:    svcURL,
		UserAgent: userAgent,
	}

	c.common.client = c
	c.Bill = (*BillService)(&c.common)
	c.CashReceipt = (*CashReceiptService)(&c.common)
	c.Coin = (*CoinService)(&c.common)
	c.CostCenter = (*CostCenterService)(&c.common)
	c.CryptoTransaction = (*CryptoTransactionService)(&c.common)
	c.Currency = (*CurrencyService)(&c.common)
	c.Customer = (*CustomerService)(&c.common)
	c.Invoice = (*InvoiceService)(&c.common)
	c.Item = (*ItemService)(&c.common)
	c.Job = (*JobService)(&c.common)
	c.Journal = (*JournalService)(&c.common)
	c.LedgerAccount = (*LedgerAccountService)(&c.common)
	c.Location = (*LocationService)(&c.common)
	c.Settings = (*SettingsService)(&c.common)
	c.Status = (*StatusService)(&c.common)
	c.Template = (*TemplateService)(&c.common)
	c.Vendor = (*VendorService)(&c.common)
	c.Warehouse = (*WarehouseService)(&c.common)
	c.Wallet = (*WalletService)(&c.common)

	return c
}

func (c *Client) NewSvcRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.SvcURL.Path, "/") {
		return nil, fmt.Errorf("SvcURL must have a trailing slash, but %q does not", c.SvcURL)
	}
	u, err := c.SvcURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse("api" + urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

type Response struct {
	*http.Response
	//nothing else yet
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		//return for context error
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		//check if there was an error with the URL
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil //empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

//need to update this error stuff to be relevant
type ErrorResponse struct {
	Response *http.Response
	Message  string  `json:"message"`
	Errors   []Error `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Errors)
}

type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
	Message  string `json:"message"`  // Message describing the error. Errors with Code == "custom" will always have this set.
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
