package main

import "fmt"

func send(c chan string, s string) {
	c <- s
}

func receive(c chan string) string {
	data := <-c
	return data
}

func main() {
	c := make(chan string, 9)
	defer close(c)
	s := "hello world!"
	go send(c, s)
	data := receive(c)
	fmt.Printf("接收到的字符串为：%s", data)
}
