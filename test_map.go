package main

import "fmt"

func main() {
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["name"] = "wxx"
	m1["city"] = "beijing"

	fmt.Printf("m1: %v\n", m1)
	for key, value := range m1 {
		fmt.Printf("%s: %s\n", key, value)
	}
}
