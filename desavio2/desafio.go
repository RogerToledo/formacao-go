package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		switch{
		case (i % 3 == 0) && (i % 5 == 0):
			fmt.Println("PIN-PAN")
		case i % 3 == 0:
			fmt.Println("PIN")
		case i % 5 == 0:
			fmt.Println("PAN")
		default:
			fmt.Println(i)
		}
	}  
}
