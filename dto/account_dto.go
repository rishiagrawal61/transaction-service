package dto

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type FetchAccountDetailsRequest struct {
	AccountID string `json:"accountId"`
}

type AccountResponse struct {
	AccountID      string `json:"id"`
	DocumentNumber string `json:"document_number"`
	CreatedAt      string `json:"created_at"`
}
