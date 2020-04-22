package service

import (
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

type Service interface {
	Balance() (*models.Balance, error)

	TaskByID(taskID uint64) (*models.Task, error)
	TaskByCustomerTaskID(customerTaskID string) (*models.Task, error)

	WorkerList() ([]models.Worker, error)

	PaymentOnTask(task contracts.PaymentRequest) (*models.Payment, error)
}

type Repository interface {
	//Info
	GetBalance() (*models.Balance, error)

	//Task
	GetTaskByID(taskID uint64) (*models.Task, error)
	GetTaskByMerchantTransaction(customerTaskID string) (*models.Task, error)

	//Worker
	GetWorkerList() ([]models.Worker, error)

	//Payment
	GetPaymentOnTask(task contracts.PaymentRequest) (*models.Payment, error)
}
