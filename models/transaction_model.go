package models

import "time"

type Transaction struct {
	TransactionId     int64     `gorm:"primaryKey" json:"id"`
	AccountId         string    `gorm:"type:varchar(100);not null" json:"account_id"`
	TransactionTypeId string    `json:"transaction_type_id"`
	Amount            float64   `json:"amount"`
	CreatedAt         time.Time `json:"created_at"`
}
