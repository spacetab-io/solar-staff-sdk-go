package models

type PaymentData struct {
	Amount   float64      `json:"payout_amount"`
	Currency CurrencyCode `json:"payout_currency"`
	Exchange int          `json:"payout_exchange"`
}

type RealPaymentDetails struct {
	Amount           float64
	CommissionAmount float64
}

type RealPaymentData struct {
	Amount   float64            `json:"real_amount"`
	Currency CurrencyCode       `json:"real_currency"`
	Exchange int                `json:"real_exchange"`
	Details  RealPaymentDetails `json:"real_details"`
}
