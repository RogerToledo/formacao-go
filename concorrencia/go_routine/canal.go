package concorrencia

import (
	"fmt"
	"time"
)

func Ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
	
}
