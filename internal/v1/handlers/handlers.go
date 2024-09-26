package handlers

import (
	"log"
	httphelpers "seguridad-cicd/pkg/http"
	"time"
)

func ExecuteHandlers(client httphelpers.HTTPClient, seguridadURL string) {

	MakeRequest(client, "GET", seguridadURL+"/health", nil)

	MakeRequest(client, "GET", seguridadURL+"/alumnos", nil)

	MakeRequest(client, "POST", seguridadURL+"/alumno", []byte{})

	MakeRequest(client, "POST", seguridadURL+"/accounts/create", []byte{})

	time.Sleep(20 * time.Second)
}

func MakeRequest(client httphelpers.HTTPClient, method, url string, payload []byte) {
	response, statusCode, err := httphelpers.Request(client, payload, url, method)
	if err != nil {
		log.Printf("Error making %s request to %s: %v", method, url, err)
		return
	}
	log.Printf("%s response from %s: %s (Status: %d)", method, url, string(response), statusCode)
}
