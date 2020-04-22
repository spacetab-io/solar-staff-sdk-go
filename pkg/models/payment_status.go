package models

type PaymentStatus struct {
	ID   uint64 `json:"status_id"`
	Name string `json:"status"`
}
