package validator

import (
	"strings"
	"transaction-service/dto"
)

func ValidateAccountCreateRequest(req dto.CreateAccountRequest) map[string]string {
	errs := map[string]string{}

	if strings.TrimSpace(req.DocumentNumber) == "" {
		errs["document_number"] = "Document number is required"
	}
	if len(req.DocumentNumber) > 20 {
		errs["document_number"] = "Document number cannot exceed 20 characters"
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ValidateAccountDetailsFetchRequest(req dto.FetchAccountDetailsRequest) map[string]string {
	errs := map[string]string{}

	if strings.TrimSpace(req.AccountID) == "" {
		errs["account_id"] = "Account ID is required"
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
