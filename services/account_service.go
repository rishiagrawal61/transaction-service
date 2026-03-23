package services

import (
	"fmt"
	"strconv"
	"transaction-service/dto"
	"transaction-service/models"
	"transaction-service/repository"
)

type AccountService interface {
	CreateAccount(req dto.CreateAccountRequest) (dto.AccountResponse, error)
	FetchAccountDetails(req dto.FetchAccountDetailsRequest) (dto.AccountResponse, error)
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(req dto.CreateAccountRequest) (dto.AccountResponse, error) {
	account := models.Account{
		DocumentNumber: req.DocumentNumber,
	}
	_, err := s.repo.FindByDocumentNumber(req.DocumentNumber)
	if err == nil {
		return getAccountResponse(account), fmt.Errorf("account with document number %s already exists", req.DocumentNumber)
	}
	account, err = s.repo.Insert(account)
	if err != nil {
		return dto.AccountResponse{}, err
	}
	return getAccountResponse(account), nil
}

func (s *accountService) FetchAccountDetails(req dto.FetchAccountDetailsRequest) (dto.AccountResponse, error) {
	account, err := s.repo.FindByID(req.AccountID)
	if err != nil {
		return dto.AccountResponse{}, err
	}
	return getAccountResponse(account), err
}

func getAccountResponse(account models.Account) dto.AccountResponse {
	return dto.AccountResponse{AccountID: strconv.FormatInt(account.ID, 10), DocumentNumber: account.DocumentNumber, CreatedAt: account.CreatedAt.String()}
}
