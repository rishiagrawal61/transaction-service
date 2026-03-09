package models

type TransactionType struct {
	TransactionType_ID int    `gorm:"primaryKey" json:"id"`
	Description        string `gorm:"type:varchar(100);not null" json:"description"`
	CreatedAt          string `json:"created_at"`
}
