package main

import (
	"net"
	"time"
)

func main() {
	udpServer, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, _ := net.DialUDP("udp", nil, udpServer)
	defer conn.Close()
	conn.Write([]byte("Hello UDP Server"))
	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		println(string(buf))
		time.Sleep(time.Second)
	}
}
