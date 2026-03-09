package routes

import (
	"log"
	"net/http"
	"transaction-service/config"
	"transaction-service/container"
	"transaction-service/middleware"
)

func NewRouter(cfg config.Config, c *container.Container) http.Handler {
	mux := http.NewServeMux()

	if cfg.EnableRestApi {
		log.Println("Starting REST API server...")
		RegisterApi(mux, c)
	}

	return middleware.Logger(mux)
}
