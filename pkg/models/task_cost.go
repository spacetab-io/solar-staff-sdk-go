package models

type Fee struct {
	Amount  float64 `json:"amountOfCommission"`
	Percent float64 `json:"percentOfCommission"`
}

type Currency struct {
	Title string
	Code  string `json:"short_title"`
}

type Cost struct {
	Price float64 `json:"amount"`
	Fee
	FullPrice float64
	Currency
}
