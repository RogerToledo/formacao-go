package main

import (
	"dio/formacao-go/manipulando-json/api_cadastro/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rota := mux.NewRouter()
	rota.HandleFunc("/cliente", handler.ClienteHandler).Methods("POST")

	fmt.Println("Escutando na porta : 8081")

	log.Fatal(http.ListenAndServe(":8081", rota))
}
