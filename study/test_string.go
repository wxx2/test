package main

import "fmt"

func main() {
	var s string = "hello world\\"
	var ss string = `hello world ||\\!!$%^&*(
						hello world`
	fmt.Println(s)
	fmt.Println(ss)
}
