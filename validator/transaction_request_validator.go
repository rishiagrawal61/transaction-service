package validator

import (
	"strconv"
	"strings"
	"transaction-service/dto"
	"transaction-service/repository"
)

type TransactionValidator struct {
	transactionTypeRepo repository.TransactionTypeRepository
}

func NewTransactionValidator(transactionTypeRepo repository.TransactionTypeRepository) *TransactionValidator {
	return &TransactionValidator{transactionTypeRepo: transactionTypeRepo}
}

func (t *TransactionValidator) ValidateTransactionRequest(req dto.TransactionCreateDTO) map[string]string {
	errs := map[string]string{}

	if req.Amount <= 0 {
		errs["amount"] = "Amount must be greater than zero"
	}
	if req.Amount > 5000000 {
		errs["amount"] = "Amount must not exceed 50,00,000 lacs"
	}
	if strings.TrimSpace(req.AccountID) == "" {
		errs["account_id"] = "Account ID is required"
	}
	if strings.TrimSpace(req.TransactionType) == "" {
		errs["operation_type_id"] = "Operation type is required"
	}
	transactionTypeInt, err := strconv.Atoi(req.TransactionType)
	if err != nil {
		errs["operation_type_id"] = "Operation type must be a valid number"
	} else {
		storedTransactionTypes, err := t.transactionTypeRepo.GetTypes()
		if err != nil {
			errs["operation_type_id"] = "Error fetching Operation types"
		}
		if _, ok := storedTransactionTypes[transactionTypeInt]; !ok {
			errs["operation_type_id"] = "Invalid Operation type"
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
