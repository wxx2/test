package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 3)

	t1 := time.Now()
	fmt.Printf("Nowtime is %v\n", t1)
	<-timer1.C
	t2 := time.Now()
	fmt.Printf("Nowtime is %v\n", t2)
}
