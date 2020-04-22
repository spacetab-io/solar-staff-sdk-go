package service

import (
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/contracts"
	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

func (s *service) PaymentOnTask(task contracts.PaymentRequest) (*models.Payment, error) {
	return s.r.GetPaymentOnTask(task)
}
