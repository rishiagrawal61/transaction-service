package models

import "time"

type Account struct {
	ID             int64     `gorm:"primaryKey" json:"id"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
}
