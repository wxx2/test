package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nanotime := int64(time.Now().Nanosecond())
	rand.Seed(nanotime)

	for i := 0; i < 250; i++ {
		num1 := rand.Intn(500)
		num2 := rand.Intn(2)
		var num int
		if num2 == 0 {
			num = num1
		} else {
			num = num1 * (-1)
		}
		fmt.Printf("%d,", num)
	}
}
