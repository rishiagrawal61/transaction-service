package dto

type TransactionCreateDTO struct {
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"operation_type_id"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"id"`
	AccountId       string  `json:"document_number"`
	TransactionType string  `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	CreatedAt       string  `json:"created_at"`
}
