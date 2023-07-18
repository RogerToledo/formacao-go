package main

import (
	"log"
	"net/http"
)

func main() {
	porta := ":3000"
	fs := http.FileServer(http.Dir("./estatic"))
	http.Handle("/", fs)
	log.Println("Listening", porta)
	err := http.ListenAndServe(porta, nil)
	if err != nil {
		log.Fatal(err)
	}
}