package softledger

type LedgerAccount struct {
	ID                      *int64  `json:"_id"`
	Name                    *string `json:"name"`
	NaturalBalance          *string `json:"naturalBalance"`
	Description             *string `json:"description"`
	Number                  *string `json:"number"`
	Type                    *string `json:"type"`
	Subtype                 *string `json:"subtype"`
	IncludeLocationChildren *bool   `json:"includeLocationChildren"`
	LocationId              *int64
	Location                *Location
}
