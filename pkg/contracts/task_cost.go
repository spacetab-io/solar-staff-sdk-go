package contracts

import (
	"strconv"

	"github.com/spacetab-io/solar-staff-sdk-go/pkg/models"
)

type Fee struct {
	Amount  string `json:"amountOfCommission"`
	Percent string `json:"percentOfCommission"`
}

func (f Fee) ToModel() (*models.Fee, error) {
	amount, err := strconv.ParseFloat(f.Amount, 64)
	if err != nil {
		return nil, err
	}

	percent, err := strconv.ParseFloat(f.Percent, 64)
	if err != nil {
		return nil, err
	}

	return &models.Fee{
		Amount:  amount,
		Percent: percent,
	}, nil
}

type Cost struct {
	Price string `json:"amount"`
	Fee
	FullPrice string
	models.Currency
}

func (c Cost) ToModel() (*models.Cost, error) {
	price, err := strconv.ParseFloat(c.Price, 64)
	if err != nil {
		return nil, err
	}
	fullPrice, err := strconv.ParseFloat(c.FullPrice, 64)
	if err != nil {
		return nil, err
	}
	fee, err := c.Fee.ToModel()

	return &models.Cost{
		Price:     price,
		Fee:       *fee,
		FullPrice: fullPrice,
		Currency:  c.Currency,
	}, nil
}
