package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transaction-service/dto"
	"transaction-service/services"
	"transaction-service/validator"
)

type TransactionHandler struct {
	service   services.TransactionService
	validator *validator.TransactionValidator
}

func NewTransactionHandler(service services.TransactionService, validator *validator.TransactionValidator) *TransactionHandler {
	return &TransactionHandler{service: service, validator: validator}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Errorf("handler.CreateTransaction, invalid request body: %w", err))
	}

	if errs := h.validator.ValidateTransactionRequest(req); len(errs) > 0 {
		writeValidationErrors(w, errs)
		return
	}

	transaction, err := h.service.CreateTransaction(req)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Errorf("handler.CreateTransaction, unable to create txn: %w", err))
		return
	}

	log.Printf("Transaction created successfully: %v", transaction)
	writeJSON(w, http.StatusCreated, transaction)
}
