package softledger

type Location struct {
	ID          *string     `json:"_id"`
	Id          *string     `json:"id"`
	Name        *string     `json:"name"`
	Currency    *string     `json:"currency"`
	Description *string     `json:"description"`
	Parent_id   *int64      `json:"parent_id"`
	Children    []*Location `json:"children"`
}
