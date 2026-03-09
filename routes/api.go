package routes

import (
	"net/http"
	"transaction-service/container"
)

func RegisterApi(mux *http.ServeMux, c *container.Container) {
	accounts := c.AccountHandler()
	transactions := c.TransactionHandler()

	mux.HandleFunc("POST /accounts", accounts.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", accounts.GetAccountByID)
	mux.HandleFunc("POST /transactions", transactions.CreateTransaction)
}
