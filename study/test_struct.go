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

//go:noinline
func (m Man) f1() {
	fmt.Println(m.name)
	fmt.Println(m.Age)
	m.name = "ttt"
}

//go:noinline
func (p *Man) f2() {
	fmt.Println(p.name)
	p.name = "edy"
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
	//fmt.Println(ptr)
	//fmt.Println(*ptr)
	//fmt.Println(&ptr)
	//fmt.Println(ptr.name)

	//var ptr1 *Man = &wxx
	//fmt.Println(ptr1)
	//fmt.Println(*ptr1)
	//fmt.Println(&ptr1)
	//fmt.Println(ptr1.name)

	wxx.f1()
	fmt.Println(ptr.name)
	wxx.f2()
	fmt.Println(ptr.name)
}
