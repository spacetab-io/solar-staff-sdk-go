package models

type Payment struct {
	WorkerEmail     string `json:"email,omitempty"`
	WorkerID        uint64 `json:"worker_id"`
	TaskID          uint64 `json:"task_id,omitempty"`
	CustomerTaskID  string `json:"merchant_transaction"`
	CardID          uint64 `json:"card_id,omitempty"`
	Currency        CurrencyCode
	Amount          float64
	TodoTypeID      uint64 `json:"todo_type"`
	TodoAttribute   string `json:"todo_attributes"`
	TaskTitle       string
	TaskDescription string
}
