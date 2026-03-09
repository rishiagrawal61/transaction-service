package container

import (
	"database/sql"

	"transaction-service/handlers"
	"transaction-service/repository"
	"transaction-service/services"
	"transaction-service/validator"
)

type Container struct {
	DB *sql.DB
}

func New(DB *sql.DB) *Container {
	return &Container{
		DB: DB,
	}
}

func (c *Container) AccountHandler() *handlers.AccountHandler {
	repo := repository.NewAccountRepository(c.DB)
	svc := services.NewAccountService(repo)
	return handlers.NewAccountHandler(svc)
}

func (c *Container) TransactionHandler() *handlers.TransactionHandler {
	repo := repository.NewTransactionRepository(c.DB)
	accountRepo := repository.NewAccountRepository(c.DB)
	transactionTypeRepo := repository.NewTransactionTypeRepository(c.DB)
	svc := services.NewTransactionService(repo, accountRepo)
	validator := validator.NewTransactionValidator(transactionTypeRepo)
	return handlers.NewTransactionHandler(svc, validator)
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
