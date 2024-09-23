package main

import (
	"fmt"
	"log"
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
	case "/api/v1":
		ctx.SetContentType("text/plain; charset=utf-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("API v1")
	case "/api/v2":
		ctx.SetContentType("text/plain; charset=utf-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("API v2")
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
	go startServer()
	go Tests()
	fmt.Println("Servidor iniciado. Ejecuta `test.go` para validar los endpoints.")
	time.Sleep(time.Hour)
}
