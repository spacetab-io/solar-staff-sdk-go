package contracts

import (
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

type PaymentRequest struct {
	WorkerEmail     string `json:"email,omitempty"`
	WorkerID        uint64 `json:"worker_id"`
	TaskID          uint64 `json:"task_id,omitempty"`
	CustomerTaskID  string `json:"merchant_transaction"`
	CardID          uint64 `json:"card_id,omitempty"`
	Currency        models.CurrencyCode
	Amount          float64
	TaskCategoryID  uint64 `json:"todo_type,omitempty"`
	ProjectName     string `json:"todo_attributes,omitempty"`
	TaskTitle       string
	TaskDescription string
}

func (r PaymentRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"action":               "payout",
		"worker_id":            r.WorkerID,
		"merchant_transaction": r.CustomerTaskID,
		"currency":             r.Currency,
		"amount":               r.Amount,
		"task_title":           r.TaskTitle,
		"task_description":     r.TaskDescription,
		"todo_type":            r.TaskCategoryID,
		"todo_attributes":      r.ProjectName,
	}
}

type PaymentResponse struct {
	Currency       models.CurrencyCode
	TransactionID  uint64
	TaskID         uint64
	CardID         uint64
	CustomerTaskID string `json:"merchant_transaction"`
	models.PaymentStatus
	models.PaymentData
	models.RealPaymentData
	PaidAt *models.SSTime
}
