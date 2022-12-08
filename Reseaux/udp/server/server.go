package main

import (
	"net"
	"time"
)

func main() {
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 8080})
	defer listener.Close()
	buf := make([]byte, 1024)
	for {
		_, addr, _ := listener.ReadFromUDP(buf)
		go manageClient(listener, addr)
	}
}

func manageClient(server *net.UDPConn, addr *net.UDPAddr) {
	for {
		now := []byte(time.Now().Format(time.ANSIC))
		server.WriteToUDP(now, addr)
		time.Sleep(time.Second)
	}
}
