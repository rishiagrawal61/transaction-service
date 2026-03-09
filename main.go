package main

import (
	"log"
	"net/http"
	"transaction-service/config"
	"transaction-service/container"
	"transaction-service/db"
	"transaction-service/routes"
)

func main() {
	config := config.Load()

	DB, err := db.Connect(db.Config(config.DB))
	if err != nil {
		log.Fatalf("[db] primary connection failed: %v", err)
	}

	c := container.New(DB)
	defer c.Close()

	router := routes.NewRouter(config, c)

	port := config.Port
	log.Printf("Server listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
