package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

type Test []byte

func (test *Test) test_print() {
	fmt.Println("hello word")
}
func main() {
	fmt.Println("start...")

	addr := net.JoinHostPort("172.17.6.57", "22")
	fmt.Println(addr)

	name, _ := net.LookupAddr("172.17.6.57")
	fmt.Println(name)

	addrs, _ := net.InterfaceAddrs()
	fmt.Println(addrs)

	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	test := Test{}
	test.test_print()
	fmt.Println(reflect.TypeOf(test))

	var t []byte
	fmt.Println(reflect.TypeOf(t))

	fmt.Println("end...")
}
