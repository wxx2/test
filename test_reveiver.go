package main

import (
	"fmt"
)

type man struct {
	name string
	age  int
}

func getname(m man) {
	fmt.Println(m.name)
}

func getname2(p *man) {
	fmt.Println(p.name)
	// 自动解引用
	fmt.Println((*p).name)
}

func main() {
	var wxx man
	wxx = man{
		"wxx",
		18,
	}
	getname(wxx)
	getname2(&wxx)
}
