package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LABEL1
		}
		fmt.Println(i)
	}

LABEL1:
	fmt.Println("end...")
}
