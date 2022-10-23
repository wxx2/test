package main

import (
	"fmt"
	"net"
)

func main() {

	//连接给定的network地址
	address := "127.0.0.1:8080"
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Println(fmt.Sprintf("连接失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("连接成功, address:%s", address))

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭连接失败, err:%s", err))
		}
		fmt.Println(fmt.Sprintf("关闭连接成功"))
	}(conn)

	msg := fmt.Sprintf("123456789")
	msgByte := []byte(msg)
	n, err := conn.Write(msgByte)
	if err != nil {
		fmt.Println(fmt.Sprintf("client发送失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("client发送成功%d个字节", n))
}
