package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		ctx.SetContentType("text/plain; charset=utf-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("Hola, Mundo!")
	case "/hello":
		ctx.SetContentType("application/json; charset=utf-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"message": "Hello, world!"}`)
	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString("Ruta no encontrada")
	}
}

func startServer() {
	fmt.Println("Servidor escuchando en el puerto 8080...")
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s\n", err)
	}
}

func main() {
	go startServer() // Inicia el servidor en un goroutine

	// Espera para que el servidor arranque
	fmt.Println("Esperando 10 segundos para que el servidor se inicie...")
	time.Sleep(10 * time.Second)

	// Realiza las pruebas
	testEndpoints()
}

func testEndpoints() {
	// Prueba para el endpoint "/"
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		log.Fatalf("Error al hacer solicitud: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || string(body) != "Hola, Mundo!" {
		log.Fatalf("Error en el endpoint /: Código de estado %v, Cuerpo: %s", resp.StatusCode, body)
	}

	// Prueba para el endpoint "/hello"
	resp, err = http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatalf("Error al hacer solicitud: %v", err)
	}
	body, _ = io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || string(body) != `{"message": "Hello, world!"}` {
		log.Fatalf("Error en el endpoint /hello: Código de estado %v, Cuerpo: %s", resp.StatusCode, body)
	}

	fmt.Println("Todas las pruebas pasaron correctamente.")
}
