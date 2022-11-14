package main

import "fmt"

type Animal interface {
	Sleep()
}

type Bird interface {
	Fly()
}

type Pigeon interface {
	Animal
	Bird
}

type PIGEON struct {
	Name string
	Age  int
}

func (p *PIGEON) Fly() {
	fmt.Printf("%v can fly...\n", p.Name)
}

//func (p *PIGEON) Sleep() {
//	fmt.Printf("%v can sleep...\n", p.Name)
//}

func main() {
	gg := PIGEON{
		Name: "ggg",
		Age:  18,
	}

	gg.Fly()
	//gg.Sleep()
}
