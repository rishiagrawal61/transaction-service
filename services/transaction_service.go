package services

import (
	"strconv"
	"transaction-service/dto"
	"transaction-service/models"
	"transaction-service/repository"
	"transaction-service/strategies"
)

type TransactionService interface {
	CreateTransaction(req dto.TransactionCreateDTO) (dto.TransactionResponse, error)
}

type transactionService struct {
	repo        repository.TransactionRepository
	accountRepo repository.AccountRepository
}

func NewTransactionService(repo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionService {
	return &transactionService{repo: repo, accountRepo: accountRepo}
}

func (s *transactionService) CreateTransaction(req dto.TransactionCreateDTO) (dto.TransactionResponse, error) {
	_, err := s.accountRepo.FindByID(req.AccountID)
	if err != nil {
		return dto.TransactionResponse{}, err
	}
	transaction, err := s.resolveTransactionStrategy(req).Create(s.repo)
	if err != nil {
		return dto.TransactionResponse{}, err
	}
	return getTransactionResponse(transaction), nil
}

func (t *transactionService) resolveTransactionStrategy(req dto.TransactionCreateDTO) strategies.TransactionStrategy {
	switch req.TransactionType {
	case "1":
		return strategies.NormalPurchaseStrategy{TransactionDTO: req}
	case "2":
		return strategies.PurchageWithInstallmentStrategy{TransactionDTO: req}
	case "3":
		return strategies.WithdrawalStrategy{TransactionDTO: req}
	case "4":
		return strategies.CreditVoucherStrategy{TransactionDTO: req}
	default:
		return nil
	}
}

func getTransactionResponse(transaction models.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   strconv.FormatInt(transaction.TransactionId, 10),
		AccountId:       transaction.AccountId,
		TransactionType: transaction.TransactionTypeId,
		Amount:          transaction.Amount,
		CreatedAt:       transaction.CreatedAt.String(),
	}
}
