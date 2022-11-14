package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func addd() {
	atomic.AddInt32(&i, 1)
}

func subb() {
	atomic.AddInt32(&i, -1)
}

func main() {
	for k := 0; k < 10; k++ {
		for j := 0; j < 100; j++ {
			go addd()
			go subb()
		}
		time.Sleep(time.Second * 1)
		fmt.Printf("i's value is %v\n", i)
	}
}
