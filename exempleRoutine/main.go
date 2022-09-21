package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	time.Sleep(time.Millisecond * 500)
	go ponger(c)
	var input string
	fmt.Scanln(&input)
}

func pinger(c chan<- string) {
	for {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * 1)
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}
