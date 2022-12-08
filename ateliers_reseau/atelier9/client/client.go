package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Print(err)
	}
	defer dial.Close()

	for {
		_, err2 := io.WriteString(dial, "Hello")
		if err2 != nil {
			log.Print(err)
		}
		time.Sleep(time.Second)
	}
}
