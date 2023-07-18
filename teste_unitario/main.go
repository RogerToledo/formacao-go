package main

import (
	"dio/formacao-go/manipulando-json/teste_unitario/calculadora"
	"fmt"
)

func main() {
	numeros := []int {1, 2, 3}
	soma := calculadora.Soma(numeros)
	fmt.Println("Soma: ", soma)

	sub := calculadora.Subtracao(numeros)
	fmt.Println("Subtração: ", sub)

	mult := calculadora.Multiplicacao(numeros)
	fmt.Println("Multiplicacao: ", mult)

	numerosD := []int{10, 2, 3}
	div, erro := calculadora.Divisao(numerosD)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("Divisao: ", div)
	}

}
