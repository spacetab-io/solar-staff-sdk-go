package models

type Balance struct {
	Rub     float64
	Details BalanceDetails
}

type BalanceDetails struct {
	Cards    float64
	Qiwi     float64
	Webmoney float64
}
