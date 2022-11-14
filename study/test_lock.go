package main

import (
	"os"
	"sync"
)

var lock sync.RWMutex

const path = "study/test.txt"

func writefilestring(path string) {
	lock.RLock()
	f, _ := os.OpenFile(path, os.O_RDWR, 0775)
	f.WriteString("hello world")
	lock.RUnlock()
}
func writefileint(path string) {
	f, _ := os.OpenFile(path, os.O_RDWR, 0775)
	f.WriteString(string("123456"))
}

func main() {
	writefilestring(path)
	writefileint(path)
}
