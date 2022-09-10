package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Ref       string  `json:"ref"`
	Date      string  `json:"date"`
	Amount    float64 `json:"amount"`
	AccountId uint    `json:"account_id"`
	Account   Account
	Reson     string  `json:"reason"`
	Status    string  `json:"status"`
	Channel   string  `json:"channel"`
	Fees      float32 `json:"fees"`
}
