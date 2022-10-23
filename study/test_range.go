package main

import (
	"fmt"
	"time"
)

func main() {
	// for range 遍历数组、切片、字符串、map、通道
	// 数组切片返回索引和值
	// map返回键和值
	// 通道只返回通道内的值
	var arr = [5]int{1, 2, 5, 4, 3}
	var m = map[string]interface{}{"name": "wxx", "age": 18}
	c := make(chan string, 10)
	go func(c chan string) {
		r := <-c
		fmt.Printf("r: %v %T", r, r)
	}(c)
	c <- "hello world!"
	defer close(c)

	for i, v := range arr {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	for k, v := range m {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
	time.Sleep(time.Second * 3)
}
