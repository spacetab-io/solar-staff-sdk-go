package service

import "github.com/spacetab-io/solar-staff-sdk-go/pkg/models"

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

type Service interface {
	Balance() (*models.Balance, error)

	TaskByID(taskID uint64) (*models.Task, error)
	TaskByMerchantTransaction(transaction string) (*models.Task, error)

	WorkerList() ([]models.Worker, error)
}

type Repository interface {
	GetBalance() (*models.Balance, error)

	GetTaskByID(taskID uint64) (*models.Task, error)
	GetTaskByMerchantTransaction(transaction string) (*models.Task, error)

	GetWorkerList() ([]models.Worker, error)
}
