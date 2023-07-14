package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Usuario struct {
	Nome     string `json: "nome"`
	CorSabre string `json: "corSabre"`
	Mestre   string `json: "mestre"`
}

type Usuarios struct {
	Usuarios []Usuario `json: "usuarios"`
}

func main() {
	jsonFile, err := os.Open("manipulando-json/usuarios.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	var usuarios Usuarios

	if err := json.Unmarshal(byteValue, &usuarios); err != nil {
		log.Fatal(err)
	}

	for _, u := range usuarios.Usuarios {
		fmt.Println("Nome:", u.Nome)
		fmt.Println("Cor do Sabre:", u.CorSabre)
		fmt.Println("Mestre:", u.Mestre)
		fmt.Println()
	}
}