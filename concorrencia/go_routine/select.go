package concorrencia

import "time"

func One(chanOne chan string) {
	time.Sleep(1 * time.Second)

	chanOne <- "um"
}

func Two(chanTwo chan string) {
	time.Sleep(1 * time.Second)

	chanTwo <- "dois"
}