package main

import "fmt"

const (
	name = "wxx"
	age  = 18
)

const school = "hohai"

func main() {
	fmt.Println(fmt.Sprintf("my name is %s, and age is %d", name, age))
	fmt.Println(fmt.Sprintf("my school is %s", school))
}
