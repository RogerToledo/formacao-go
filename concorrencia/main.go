package main

import (
	concorrencia "dio/formacao-go/concorrencia/go_routine"
	"fmt"
)

func main() {
	// go concorrencia.F(0)
	// go concorrencia.F(1)

	// var escreva string
	// fmt.Scanln(&escreva)

	even, odd := concorrencia.EvenOdd(10)

	for _, e := range even {
		fmt.Printf("Even: %d \n", e)
	}
	fmt.Println("------------------")

	for _, o := range odd {
		fmt.Printf("Odd: %d", o)
	}
	
}