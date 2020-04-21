package rest

import (
	"encoding/json"
	"fmt"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func (r *Repo) GetWorkerList() ([]models.Worker, error) {
	resp, err := r.sendPOST("workers", map[string]interface{}{"action": "workers_list"})
	if err != nil {
		return nil, err
	}

	var data struct {
		Success  bool
		Code     int
		Response []models.Worker
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return data.Response, nil
}
