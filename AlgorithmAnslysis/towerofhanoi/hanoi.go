package main

import "fmt"

var A [5]int = [5]int{1, 2, 3, 4, 5}
var B, C [5]int

func move(n int, A, C *[5]int) {
	C[n-1] = A[n-1]
}

func Hanoi(n int, A, B, C *[5]int) {
	if n == 1 {
		move(1, A, C)
		return
	}
	Hanoi(n-1, A, C, B)
	move(n-1, A, C)
	Hanoi(n-1-1, B, A, C)
}

func main() {
	Hanoi(5, &A, &B, &C)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
