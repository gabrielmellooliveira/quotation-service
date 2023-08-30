package main

import (
	"log"
	"net/http"
	"quotation-server/handlers"
)

func main() {
	mux := http.NewServeMux()

    mux.HandleFunc("/cotacao", handlers.QuotationHandler)

    log.Println("Iniciando o servidor em http://localhost:8080")

	http.ListenAndServe(":8080", mux)
}
