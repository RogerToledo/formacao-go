package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Descriptions []Description `json: "descriptions"`
}

type Description struct {
	Description string `json: "description"`
	Language Language `json: "language"`
}

type Language struct {
	Name string `json: "name"`
	Url  string `json: "url"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resp Response

	if err := json.Unmarshal(responseData, &resp); err != nil {
		log.Fatal(err)
	}

	for _, r := range resp.Descriptions {
		fmt.Println("Description: ", r.Description)
		fmt.Println("Name: ", r.Language.Name)
		fmt.Println("URL: ", r.Language.Url)
		fmt.Println()
	}

}
