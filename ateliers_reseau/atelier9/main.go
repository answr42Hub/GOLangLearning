package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	server, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := io.ReadAll(conn)
		if err != nil {
			return
		}
		fmt.Print(string(message))
		time.Sleep(1 * time.Second)
	}
}
