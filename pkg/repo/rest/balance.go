package rest

import (
	"encoding/json"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func (r Repo) GetBalance() (*models.Balance, error) {
	resp, err := r.sendPOST("info", map[string]interface{}{"action": "balance"})
	if err != nil {
		return nil, err
	}

	var data struct {
		Success  bool
		Code     int
		Response contracts.BalanceResponse
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data.Response.ToModel()
}
