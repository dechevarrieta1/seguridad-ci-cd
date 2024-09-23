package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Tests() {
	// Esperar a que el servidor esté listo
	time.Sleep(5 * time.Second)

	endpoints := []struct {
		name     string
		path     string
		expected string
	}{
		{"Raíz", "http://localhost:8080/", "Hola, Mundo!"},
		{"Hello", "http://localhost:8080/hello", `{"message": "Hello, world!"}`},
		{"No encontrado", "http://localhost:8080/no-existe", "Ruta no encontrada"},
		{"Nuevo endpoint 1", "http://localhost:8080/api/v1", "API v1"},
		{"Nuevo endpoint 2", "http://localhost:8080/api/v2", "API v2"},
	}

	failedTests := 0

	for _, e := range endpoints {
		fmt.Printf("Probando endpoint %s: %s\n", e.name, e.path)
		resp, err := http.Get(e.path)
		if err != nil {
			log.Printf("Error al hacer la petición a %s: %v\n", e.path, err)
			failedTests++
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error al leer la respuesta de %s: %v\n", e.path, err)
			failedTests++
			continue
		}

		if string(body) != e.expected {
			log.Printf("La respuesta de %s no coincide. Esperado: %s, Obtenido: %s\n", e.path, e.expected, string(body))
			failedTests++
		} else {
			fmt.Printf("Prueba exitosa para %s\n", e.name)
		}
	}

	if failedTests > 0 {
		fmt.Printf("Fallaron %d pruebas\n", failedTests)
		os.Exit(1)
	} else {
		fmt.Println("Todas las pruebas pasaron exitosamente")
	}
}
