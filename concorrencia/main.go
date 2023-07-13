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

	// ----------------------------------------
	c := make(chan string)

	go concorrencia.Ping(c)
	go concorrencia.Imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

	// ----------------------------------------

	// c1 := make(chan string)
	// c2 := make(chan string)

	// go concorrencia.One(c1)
	// go concorrencia.Two(c2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <- c1:
	// 		fmt.Println("Receba: ", msg1)
	// 	case msg2 := <- c2:
	// 		fmt.Println("Receba: ", msg2)
	// 	}
	// }

}
