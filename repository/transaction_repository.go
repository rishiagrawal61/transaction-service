package repository

import (
	"database/sql"
	"transaction-service/models"
)

type TransactionRepository interface {
	Insert(models.Transaction) (models.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Insert(transaction models.Transaction) (models.Transaction, error) {
	result, err := r.db.Exec("INSERT INTO transactions (account_id, transaction_type_id, amount) VALUES (?, ?, ?)", transaction.AccountId, transaction.TransactionTypeId, transaction.Amount)
	if err != nil {
		return models.Transaction{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Transaction{}, err
	}
	err = r.db.QueryRow("SELECT created_at FROM transactions WHERE id = ?", id).Scan(&transaction.CreatedAt)
	if err != nil {
		return models.Transaction{}, err
	}
	transaction.TransactionId = id
	return transaction, nil
}
