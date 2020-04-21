package service

import "github.com/spacetab-io/solar-staff-sdk-go/pkg/models"

func (s *service) WorkerList() ([]models.Worker, error) {
	return s.r.GetWorkerList()
}
