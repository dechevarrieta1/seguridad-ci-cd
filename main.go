package main

import (
	"net/http"
	"os"
	"seguridad-cicd/internal/v1/handlers"
	httphelpers "seguridad-cicd/pkg/http"
)

var seguridadURL = os.Getenv("SEGURIDAD_URL")

func main() {
	var client httphelpers.HTTPClient = &http.Client{}

	handlers.ExecuteHandlers(client, seguridadURL)
}
