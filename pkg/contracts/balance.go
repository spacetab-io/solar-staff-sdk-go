package contracts

import (
	"strconv"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

type BalanceResponse struct {
	Rub     string
	Details struct {
		Cards    string
		Qiwi     string
		Webmoney string
	}
}

func (b BalanceResponse) ToModel() (*models.Balance, error) {
	rub, err := strconv.ParseFloat(b.Rub, 64)
	if err != nil {
		return nil, err
	}

	cards, err := strconv.ParseFloat(b.Details.Cards, 64)
	if err != nil {
		return nil, err
	}

	qiwi, err := strconv.ParseFloat(b.Details.Qiwi, 64)
	if err != nil {
		return nil, err
	}

	webmoney, err := strconv.ParseFloat(b.Details.Webmoney, 64)
	if err != nil {
		return nil, err
	}

	return &models.Balance{
		Rub: rub,
		Details: models.BalanceDetails{
			Cards:    cards,
			Qiwi:     qiwi,
			Webmoney: webmoney,
		},
	}, nil
}
