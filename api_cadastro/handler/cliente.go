package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Cliente struct {
	Nome   string `json: "nome"`
	Idade  int64  `json: "idade"`
	Estado string `json: "estado"`
}

func ClienteHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var cliente Cliente

	erro := decoder.Decode(&cliente)
	if erro != nil {
		http.Error(w, "Erro ao decodificar Body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Println("Cliente: ", cliente)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cliente recebido com sucesso"))
}
