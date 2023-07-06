package main

import "fmt"

func main() {
	for i := 1; i <= 12; i++ {
		switch i {
		case 1:
			fmt.Printf("Janeiro -> %v, %v\n", i, parImpar(i))
		case 2:
			fmt.Printf("Fevereiro -> %v, %v\n", i, parImpar(i))
		case 3:
			fmt.Printf("Março -> %v, %v\n", i, parImpar(i))
		case 4:
			fmt.Printf("Abril -> %v, %v\n", i, parImpar(i))
		case 5:
			fmt.Printf("Maio -> %v, %v\n", i, parImpar(i))
		case 6:
			fmt.Printf("Junho -> %v, %v\n", i, parImpar(i))
		case 7:
			fmt.Printf("Julho -> %v, %v\n", i, parImpar(i))
		case 8:
			fmt.Printf("Agosto -> %v, %v\n", i, parImpar(i))
		case 9:
			fmt.Printf("Setembro -> %v, %v\n", i, parImpar(i))
		case 10:
			fmt.Printf("Outubro -> %v, %v\n", i, parImpar(i))
		case 11:
			fmt.Printf("Novembro -> %v, %v\n", i, parImpar(i))
		case 12:
			fmt.Printf("Dezembro -> %v, %v\n", i, parImpar(i))
		}
	}
}

func parImpar(num int) string {
	if num % 2 == 0 {
		return "Par"
	} else {
		return "Impar"
	}
}