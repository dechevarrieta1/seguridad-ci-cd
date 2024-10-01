package main

import (
	"log"
	"net/http"
	"os"
	"seguridad-cicd/internal/v1/handlers"
	httphelpers "seguridad-cicd/pkg/http"
)

var seguridadURL = os.Getenv("SEGURIDAD_URL")

func main() {
	log.Println("Starting the application..1")
	var client httphelpers.HTTPClient = &http.Client{}

	handlers.ExecuteHandlers(client, seguridadURL)
}
