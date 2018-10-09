package softledger

type PurchaseOrder struct {
	ID                *int64   `json:"_id"`
	Description       *string  `json:"description"`
	IssueDate         *string  `json:"issueDate"`
	DeliveryDate      *string  `json:"deliveryDate"`
	Status            *string  `json:"status"`
	Amount            *float64 `json:"amount"`
	Currency          *string  `json:"currency"`
	Notes             *string  `json:"notes"`
	Attachments       *Attachments
	TemplateId        *int64
	VendorId          *int64
	Number            *string `json:"number"`
	WarehouseId       *int64
	Warehouse         *Warehouse
	ShippingAddressId *int64
	ShippingAddress   *Address
	LocationId        *int64
	Location          *Location
	ICLocationId      *int64
	ICLocation        *Location
	LedgerAccountId   *int64
	LedgerAccount     *LedgerAccount
	POLineItems       *POLineItems
}

type POLineItems []*POLineItem

type POLineItem struct {
	Description  *string  `json:"description"`
	Quantity     *float64 `json:"quantity"`
	CostCenterId *int64
	CostCenter   *CostCenter
	KitId        *int64
	Kit          *Kit
	JobId        *int64
	Job          *Job
	Amount       *float64 `json:"amount"`
	ItemId       *int64
	Item         *Item
	KitLineItems []*POLineItem
}
