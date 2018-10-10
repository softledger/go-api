package softledger

type Contact struct {
	ID        *int64  `json:"_id"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	IsPrimary *bool   `json:"isPrimary"`
}

type Contacts []*Contact
