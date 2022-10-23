package main

import "fmt"

func main() {
	// + - * / %
	// & | ^ << >>

	var i, ii = 12, 123
	fmt.Printf("%d / %d = %d\n", ii, i, ii/i)
	fmt.Printf("%d %% %d = %d\n", ii, i, ii%i)

	var iii = 22
	fmt.Printf("二进制表示： %b\n", iii)
	fmt.Printf("左移一位：%b, %d\n", iii<<1, iii<<1)
	fmt.Printf("右移一位：%b, %d\n", iii>>1, iii>>1)
}
