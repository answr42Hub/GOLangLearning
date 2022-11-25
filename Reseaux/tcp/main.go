package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	server, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8080})
	if err != nil {
		log.Fatalln(err)
	}

	go func(l *net.TCPListener) {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		l.Close()
		os.Exit(0)
	}(server)
	
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		defer conn.Close()
		for {
			_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
			if err != nil {
				break
			}
			time.Sleep(time.Second)
		}
	}
}
