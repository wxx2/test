package main

import "fmt"

func printHelloWorld() {
	fmt.Println("hello world")
}

func hf(id int, f func()) {
	fmt.Println(id)
	f()
}

func rf(f func()) func() {
	return f
}

func main() {
	hf(1, printHelloWorld)
	rf(printHelloWorld)()
}
