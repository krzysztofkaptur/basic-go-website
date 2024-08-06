package main

import (
	"net/http"

	handlers "github.com/krzysztofkaptur/basic-go-website/pkg/handlers"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /", handlers.HomeHandler)
	server.HandleFunc("GET /about", handlers.AboutHandler)

	http.ListenAndServe(":8080", server)
}
