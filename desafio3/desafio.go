package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)

	go ping(pingChan)
	go pong(pongChan)

	for {
		select {
		case ping := <- pingChan:
			fmt.Println(ping)
		case pong := <- pongChan:
			fmt.Println(pong)
		}
	}
	
}

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(1 * time.Second)
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(1 * time.Second)
	}
}

