package handlers

import (
	"encoding/json"
	"log"
	"seguridad-cicd/internal/v1/models"
	httphelpers "seguridad-cicd/pkg/http"
	"time"
)

func ExecuteHandlers(client httphelpers.HTTPClient, seguridadURL string) {
	alumno := models.Alumno{
		Nombre:   "Juan",
		Apellido: "Sanchez",
	}
	dataReq, _ := json.Marshal(alumno)
	MakeRequest(client, "GET", seguridadURL+"/health", nil)

	MakeRequest(client, "GET", seguridadURL+"/alumnos", nil)

	//This action returns an error
	MakeRequest(client, "POST", seguridadURL+"/alumno", dataReq)

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
