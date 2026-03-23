package repository

import (
	"database/sql"
	"errors"
	"transaction-service/models"
)

type AccountRepository interface {
	FindByID(id string) (models.Account, error)
	FindByDocumentNumber(documentNumber string) (models.Account, error)
	Insert(models.Account) (models.Account, error)
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) FindByID(id string) (models.Account, error) {
	var account models.Account
	err := r.db.QueryRow("SELECT id, document_number, created_at FROM accounts WHERE id = ?", id).Scan(&account.ID, &account.DocumentNumber, &account.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Account{}, errors.New("account not found")
	}
	return account, err
}

func (r *accountRepository) FindByDocumentNumber(documentNumber string) (models.Account, error) {
	var account models.Account
	err := r.db.QueryRow("SELECT id, document_number, created_at FROM accounts WHERE document_number = ?", documentNumber).Scan(&account.ID, &account.DocumentNumber, &account.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Account{}, errors.New("account not found")
	}
	return account, err
}

func (r *accountRepository) Insert(account models.Account) (models.Account, error) {
	result, err := r.db.Exec("INSERT INTO accounts (document_number) VALUES (?)", account.DocumentNumber)
	if err != nil {
		return models.Account{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Account{}, err
	}
	err = r.db.QueryRow("SELECT created_at FROM accounts WHERE id = ?", id).Scan(&account.CreatedAt)
	if err != nil {
		return models.Account{}, err
	}
	account.ID = id
	return account, nil
}
