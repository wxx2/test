package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayhello(i int) {
	defer wg.Add(-1) //wg.Done()
	fmt.Printf("hello world %d\n", i)

}

func main() {
	for i := 0; i < 5; i++ {
		go sayhello(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("end...")
}
