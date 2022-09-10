package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name        string `json:"name"`
	Number      string `json:"number"`
	AccountType string `json:"account_type"`
	Credit      string `json:"credit"`
	Debit       string `json:"debit"`
	UserId      string `json:"user_id"`
	User        User
	Balance     float32 `json:"balance"`
	Transactions []Transaction 
}
