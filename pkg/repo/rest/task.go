package rest

import (
	"encoding/json"
	"fmt"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func (r *Repo) GetTaskByID(taskID uint64) (*models.Task, error) {
	resp, err := r.sendPOST("task", map[string]interface{}{"action": "search", "task_id": taskID})
	if err != nil {
		return nil, err
	}

	var data struct {
		Success  bool
		Code     int
		Response *contracts.TaskResponse
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return data.Response.ToModel()
}

func (r *Repo) GetTaskByMerchantTransaction(transaction string) (*models.Task, error) {
	return nil, nil
}
