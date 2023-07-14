package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Mundo!")
}

func cabecalho(w http.ResponseWriter, r *http.Request) {
	for n, cabecalhos := range r.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", n, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalho)
	
	fmt.Println("Excutando porta :8081")
	http.ListenAndServe(":8081", nil)
}
