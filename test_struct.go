package main

import "fmt"

type Person struct {
	Age int64
}

type Man struct {
	Person
	name string
	sex  string
}

func (*Person) f1() {
	println("hello word")
}

func (*Man) f2() {
	println("hello wxx")
}

func main() {

	person := Person{
		Age: 12,
	}

	wxx := Man{
		Person: person,
		name:   "ndy",
		sex:    "man",
	}

	var ptr = &wxx
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	fmt.Println(ptr.name)

	var ptr1 *Man = &wxx
	fmt.Println(ptr1)
	fmt.Println(*ptr1)
	fmt.Println(&ptr1)
	fmt.Println(ptr1.name)
}
