package main

import "fmt"

func main() {
	var a1 = []int{1, 2, 3}
	var a2 = [2]int{1, 2}
	var a3 = [...]int{1, 2, 3, 4, 5}
	var a4 = []int{}
	var a5 = [2]int{}

	fmt.Println(len(a1))
	fmt.Println(len(a2))
	fmt.Println(len(a3))
	fmt.Println(len(a4))
	fmt.Println(len(a5))
	fmt.Println("")
	fmt.Println(cap(a1))
	fmt.Println(cap(a2))
	fmt.Println(cap(a3))
	fmt.Println(cap(a4))
	fmt.Println(cap(a5))
	fmt.Println("")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
}
