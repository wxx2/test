package main

import "fmt"

func add() func(int) int {
	var x int
	return func(y int) int {
		fmt.Printf("y: %d\n", y)
		x += y
		return x
	}
}

func main() {
	var f = add()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
}
