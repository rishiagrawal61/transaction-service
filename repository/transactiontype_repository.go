package repository

import (
	"database/sql"
	"transaction-service/models"
)

type TransactionTypeRepository interface {
	GetTypes() (map[int]string, error)
}

type transactionTypeRepository struct {
	db *sql.DB
}

func NewTransactionTypeRepository(db *sql.DB) TransactionTypeRepository {
	return &transactionTypeRepository{db: db}
}

func (r *transactionTypeRepository) GetTypes() (map[int]string, error) {
	return getTransactionType(r.db)
}

// Ideally this should be fetched from Cache,
// which is populated from DB whenever updated
// or once in a day by some job or some automated process.
func getTransactionType(db *sql.DB) (map[int]string, error) {
	rows, err := db.Query(`
		SELECT id, description
		FROM transactions_types
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappedTransactionTypes = make(map[int]string)
	for rows.Next() {
		var u models.TransactionType
		if err := rows.Scan(&u.TransactionType_ID, &u.Description); err != nil {
			return nil, err
		}
		mappedTransactionTypes[u.TransactionType_ID] = u.Description
	}

	return mappedTransactionTypes, nil
}
