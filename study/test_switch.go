package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println(i)
		case 1:
			fmt.Println(i)
		case 2:
			fmt.Println(i)
		default:
			fmt.Println("欧拉欧拉！")
		}
	}
}
