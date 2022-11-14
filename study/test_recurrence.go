package main

import "fmt"

//阶乘
func jc(n float64) float64 {
	if n == 1 {
		return n
	}
	if n < 1 {
		return -1
	}
	return n * jc(n-1)
}
func main() {
	n := jc(30)
	fmt.Println(n)
}
