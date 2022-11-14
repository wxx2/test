package main

import (
	"fmt"
	"time"
)

func showMsg(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	for gap := 10; gap > 0; gap /= 2 {
		fmt.Println(gap)
	}
}
