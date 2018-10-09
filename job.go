package softledger

type Job struct {
	ID          *int64  `json:"_id"`
	Number      *string `json:"number"`
	Name        *string `json:"name"`
	Status      *string `json:"status"`
	Description *string `json:"description"`
	AgentId     *int64  `json:"AgentId"`
	Agent       *Customer
}
