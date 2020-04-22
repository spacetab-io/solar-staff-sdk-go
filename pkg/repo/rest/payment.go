package rest

import (
	"encoding/json"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func (r *Repo) GetPaymentOnTask(task contracts.PaymentRequest) (*models.Payment, error) {
	resp, err := r.sendPOST("payment", task.ToMap())
	if err != nil {
		return nil, err
	}

	var data struct {
		Success  bool
		Code     int
		Response contracts.PaymentResponse
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
