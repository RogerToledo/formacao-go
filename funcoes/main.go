package main

import (
	"dio/formacao-go/manipulando-json/funcoes/funcoes"
	"fmt"
)

func main(){

	// Função
	// notas := []float64{8.5, 7.0, 9.3, 8.7} 
	// media := funcoes.Media(notas)
	// fmt.Println(media)

	// Closure
	// x := 0
	// numero := func() int {
	// 	x++
	// 	return x
	// }

	// fmt.Println(numero())
	// fmt.Println(numero())
	// fmt.Println(numero())

	// Fatorial
	fmt.Println(funcoes.Fatorial(4))
}
