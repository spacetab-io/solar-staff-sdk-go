package service

import "github.com/spacetab-io/solar-staff-sdk-go/pkg/models"

func (s *service) TaskByID(taskID uint64) (*models.Task, error) {
	return s.r.GetTaskByID(taskID)
}

func (s *service) TaskByMerchantTransaction(transaction string) (*models.Task, error) {
	panic("implement me")
}
