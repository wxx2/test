package main

import "fmt"

const MAX int = 3

func main() {
	var p [MAX]*int

	arr := [4]int{1, 2, 3, 4}
	p[0] = &arr[0]
	p[1] = &arr[1]
	p[2] = &arr[2]
	fmt.Printf("arr[0]的地址：%v, arr[0]= %v\n", p[0], *p[0])
	fmt.Printf("arr[1]的地址：%v, arr[1]= %v\n", p[1], *p[1])
	fmt.Printf("arr[2]的地址：%v, arr[2]= %v\n", p[2], *p[2])
	fmt.Printf("arr[3]的地址：%v, arr[3]= %v\n", &arr[3], arr[3])
}
