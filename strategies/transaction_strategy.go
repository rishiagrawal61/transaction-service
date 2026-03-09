package strategies

import (
	"transaction-service/dto"
	"transaction-service/models"
	"transaction-service/repository"
)

type TransactionStrategy interface {
	Create(repo repository.TransactionRepository) (models.Transaction, error)
}

type NormalPurchaseStrategy struct {
	TransactionDTO dto.TransactionCreateDTO
}
type PurchageWithInstallmentStrategy struct {
	TransactionDTO dto.TransactionCreateDTO
}
type WithdrawalStrategy struct {
	TransactionDTO dto.TransactionCreateDTO
}
type CreditVoucherStrategy struct {
	TransactionDTO dto.TransactionCreateDTO
}

func (s NormalPurchaseStrategy) Create(repo repository.TransactionRepository) (models.Transaction, error) {
	transaction := models.Transaction{
		AccountId:         s.TransactionDTO.AccountID,
		TransactionTypeId: s.TransactionDTO.TransactionType,
		Amount:            -s.TransactionDTO.Amount,
	}

	transaction, err := repo.Insert(transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (s PurchageWithInstallmentStrategy) Create(repo repository.TransactionRepository) (models.Transaction, error) {
	transaction := models.Transaction{
		AccountId:         s.TransactionDTO.AccountID,
		TransactionTypeId: s.TransactionDTO.TransactionType,
		Amount:            -s.TransactionDTO.Amount,
	}

	transaction, err := repo.Insert(transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (s WithdrawalStrategy) Create(repo repository.TransactionRepository) (models.Transaction, error) {
	transaction := models.Transaction{
		AccountId:         s.TransactionDTO.AccountID,
		TransactionTypeId: s.TransactionDTO.TransactionType,
		Amount:            -s.TransactionDTO.Amount,
	}

	transaction, err := repo.Insert(transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (s CreditVoucherStrategy) Create(repo repository.TransactionRepository) (models.Transaction, error) {
	transaction := models.Transaction{
		AccountId:         s.TransactionDTO.AccountID,
		TransactionTypeId: s.TransactionDTO.TransactionType,
		Amount:            s.TransactionDTO.Amount,
	}

	transaction, err := repo.Insert(transaction)
	if err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}
