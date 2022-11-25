package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	dialer, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		msg, err := bufio.NewReader(dialer).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print("-->", msg)
		time.Sleep(time.Second)
	}
	//defer dialer.Close()
}
