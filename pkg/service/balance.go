package service

import "github.com/spacetab-io/solar-staff-sdk-go/pkg/models"

func (s *service) Balance() (*models.Balance, error) {
	return s.r.GetBalance()
}
