package main

import (
	"fmt"
	"os"
)

func readfile(filename str) {
	f, err := os.OpenFile("test.txt")
	if err != nil {
		print(err)
	}
	fmt.Printf("%v", f)
}

func main() {
	fmt.Println("hello")
}
