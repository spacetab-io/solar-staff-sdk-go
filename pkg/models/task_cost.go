package models

type Fee struct {
	Amount  float64 `json:"amountOfCommission"`
	Percent float64 `json:"percentOfCommission"`
}

type Cost struct {
	Price float64 `json:"amount"`
	Fee
	FullPrice float64
	Currency
}
