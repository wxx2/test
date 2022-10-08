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
	go showMsg("go")
	showMsg("down")
}
