package concorrencia

import "fmt"

func F(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}