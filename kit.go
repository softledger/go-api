package softledger

type Kit struct {
	ID          *int64  `json:"_id"`
	Number      *string `json:"number"`
	Name        *string `json:"name"`
	Sku         *string `json:"sku"`
	Description *string `json:"description"`
	ItemId      *int64  `json:"ItemId"`
	Items       []*Item
}
