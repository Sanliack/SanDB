package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:3665")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("client net.DialTCP error", err)
		return
	}
	fmt.Println(conn)
	select {}
}
