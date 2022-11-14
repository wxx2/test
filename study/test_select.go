package main

import (
	"fmt"
	"time"
)

func ssend(c1 chan int, c2 chan string) {
	c1 <- 123456
	c2 <- "hello world"
}

func main() {
	c1 := make(chan int, 9)
	c2 := make(chan string, 9)
	defer close(c1)
	defer close(c2)

	ssend(c1, c2)

	for {
		select {
		case a := <-c1:
			fmt.Printf("c1: %v\n", a)
		case a := <-c2:
			fmt.Printf("c2: %v\n", a)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
