package main

import (
	"net/http"

	"github.com/firdavsDev/go-quest/controllers"
	"github.com/firdavsDev/go-quest/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
