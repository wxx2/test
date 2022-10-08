package main

import "fmt"

func main() {
	var a1 = [10]int{1, 2, 3}
	fmt.Println(a1)
	fmt.Println("-------------------")

	for _, v := range a1 {
		fmt.Println(v)
	}
	fmt.Println("-------------------")

	for i, _ := range a1 {
		fmt.Println(i)
	}
	fmt.Println("-------------------")

	for i, v := range a1 {
		fmt.Printf("id: %d, value: %d\n", i, v)
	}
	fmt.Println("-------------------")

	fmt.Println(a1[0])
	fmt.Println(a1[2])
}
