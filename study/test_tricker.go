package main

import (
	"fmt"
	"time"
)

func main() {
	tricker := time.NewTicker(time.Second * 2)

	counter := 0
	for _ = range tricker.C {
		t := time.Now()
		fmt.Printf("tricker %d, time is %v\n", counter, t)
		counter++
		if counter > 5 {
			break
		}
	}
}
