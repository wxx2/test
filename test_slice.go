package main

import "fmt"

func main() {
	var slice1 = []int{1, 2}
	var slice2 = make([]int, 3)
	var slice3 = make([]int, 3, 5)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("-------------")
	fmt.Println(len(slice1))
	fmt.Println(len(slice2))
	fmt.Println(len(slice3))
	fmt.Println("-------------")
	fmt.Println(cap(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(cap(slice3))
	fmt.Println("-------------")

	var slice4 = []int{1, 2, 3, 6, 7, 8}
	for i, v := range slice4 {
		fmt.Printf("id: %d, value: %d\n", i, v)
	}
	fmt.Println("-------------")
	for _, v := range slice4 {
		fmt.Println(v)
	}

}
