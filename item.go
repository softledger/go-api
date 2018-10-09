package softledger

type Item struct {
	ID                   *int64   `json:"_id"`
	Number               *string  `json:"number"`
	Name                 *string  `json:"name"`
	Sku                  *string  `json:"sku"`
	Description          *string  `json:"description"`
	SalesPrice           *float64 `json:"salesPrice"`
	InvoiceAccountId     *int64   `json:"InvoiceAccountId"`
	InvoiceAccount       *LedgerAccount
	BillAccountId        *int64 `json:"BillAccountId"`
	BillAccount          *LedgerAccount
	InventoryAccountId   *int64 `json:"InventoryAccountId"`
	InventoryAccount     *LedgerAccount
	CogsAccountId        *int64 `json:"CogsAccountId"`
	CogsAccount          *LedgerAccount
	LowStockNotification *bool   `json:"lowStockNotification"`
	LowStockThreshold    *int64  `json:"lowStockThreshold"`
	LowStockEmail        *string `json:"lowStockEmail"`
	//CustomFields         *interface{} `json:"customFields"`
}
