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

type AccountHandler struct {
	service services.AccountService
}

func NewAccountHandler(service services.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Errorf("handler.CreateAccount, invalid request body: %w", err))
		return
	}

	if errs := validator.ValidateAccountCreateRequest(req); len(errs) > 0 {
		writeValidationErrors(w, errs)
		return
	}
	account, err := h.service.CreateAccount(req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("handler.CreateAccount, failed to create account: %w", err))
		return
	}

	log.Printf("Account created successfully: %s", account.AccountID)
	writeJSON(w, http.StatusCreated, account)
}

func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("accountId")
	if accountID == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("handler.GetAccountByID, invalid request body"))
		return
	}
	req := dto.FetchAccountDetailsRequest{
		AccountID: accountID,
	}

	if errs := validator.ValidateAccountDetailsFetchRequest(req); len(errs) > 0 {
		writeValidationErrors(w, errs)
		return
	}

	account, err := h.service.FetchAccountDetails(req)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Errorf("handler.GetAccountByID, failed to fetch account details: %w", err))
		return
	}

	log.Printf("Account details fetched successfully: %s", account.AccountID)
	writeJSON(w, http.StatusOK, account)
}
