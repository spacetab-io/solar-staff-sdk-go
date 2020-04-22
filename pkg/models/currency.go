package models

type CurrencyCode string

type Currency struct {
	Title string
	Code  CurrencyCode `json:"short_title"`
}
