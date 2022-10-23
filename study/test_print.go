package main

import "fmt"

func main() {
	//占位符
	// %v 相应的默认模式
	// %#v 相应值的go语法表示
	// %T 相应值的类型的go语法表示
	// %% 字面上的百分号，并非值的占位符

	var s string = "hello world"
	fmt.Printf("%v %#v %T\n", s, s, s)

	//bool
	// %t 单词true或者false
	var t bool = true
	fmt.Printf("%t %T\n", t, t)

	//整型
	// %b 二进制表示
	// %c 相应Unicode码点所表示的字符
	// %d 十进制表示
	// %o 八进制表示
	// %q 单引号围绕的字符字面值，由go语法安全的转义
	// %x 十六进制表示，字母形式为小写 a-f
	// %X 十六进制表示，字母形式为大写A-F
	// %U Unicode格式，U+1234,等同于“U+%04X”
	i := 123
	fmt.Printf("%b %c %d %o %q %x %X %U\n", i, i, i, i, i, i, i, i)

	//浮点数
	// %b 无小数部分的，指数为2的幂的科学计数法
	// %e 科学计数法
	// %E 科学计数法
	// %f 有小数点而无指数
	// %g 根据情况选择%e 活 %f以产生更紧凑的输出（无末尾0）
	// %G 根据情况选择%E 活 %F以产生更紧凑的输出（无末尾0）
	f := 123.456
	fmt.Printf("%b %e %E %.2f %.2g %.2G\n", f, f, f, f, f, f)

	//字符串与字节切片
	// %s 输出字符串表示
	// %q 双引号围绕的字符串，由go语法安全的转义
	// %x 十六进制，小写字母，每字节两个字符
	// %X 十六进制，大写字母，每字节两个字符
	var ss []byte = []byte("hello world")
	fmt.Printf("%s %q %x %X\n", ss, ss, ss, ss)

	//指针
	// %p 十六进制表示，前缀0x
	var p *int
	p = &i
	fmt.Printf("%p\n", p)
}
