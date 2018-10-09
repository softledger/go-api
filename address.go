package softledger

type Address struct {
	ID        *int64  `json:"_id"`
	Label     *string `json:"label,omitempty"`
	Line1     *string `json:"Line1,omitempty"`
	Line2     *string `json:"Line2,omitempty"`
	City      *string `json:"City,omitempty"`
	State     *string `json:"State,omitempty"`
	Zip       *string `json:"Zip,omitempty"`
	Country   *string `json:"Country,omitempty"`
	IsDefault *bool   `json:"isDefault,omitempty"`
}

type Addresses []*Address
