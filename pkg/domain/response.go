package domain

type Response struct {
	RequestId string `json:"request_id"`
	Status    int    `json:"status"`
}
