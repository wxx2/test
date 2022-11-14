package main

import "fmt"

type Human interface {
	eat()
	drink()
	love()
}

type Woman struct {
	Name string
	Age  int
}

func (w Woman) eat() {
	fmt.Printf("%v is eating...\n", w.Name)
	w.Age = 0
}

func (p *Woman) drink() {
	fmt.Printf("%v is drinking...\n", p.Name)
	p.Age = 1
}

func main() {
	xxw := Woman{
		Name: "edy",
		Age:  18,
	}

	rrr := &Woman{
		Name: "rrr",
		Age:  18,
	}

	//结构体初始化变量 + 结构体实现接口
	xxw.eat()
	fmt.Printf("%v's age is %v\n", xxw.Name, xxw.Age)
	//结构体初始化变量 + 结构体指针实现接口
	xxw.drink()
	fmt.Printf("%v's age is %v\n", xxw.Name, xxw.Age)

	//结构体指针初始化变量 + 结构体实现接口
	rrr.eat()
	fmt.Printf("%v's age is %v\n", rrr.Name, rrr.Age)
	//结构体指针初始化变量 + 结构体指针实现接口
	rrr.drink()
	fmt.Printf("%v's age is %v\n", rrr.Name, rrr.Age)
}
