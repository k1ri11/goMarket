package main

import (
	"goMarket/internal/api"
	"goMarket/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	router := api.NewRouter()

	log.Println("Starting server on port", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
