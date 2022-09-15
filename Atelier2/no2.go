package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numTwo() {
	rand.Seed(time.Now().UnixNano())
	var dist = make(map[int]int)
	var num int
	for i := 0; i < 100; i++ {
		num = rand.Intn(199) + 1
		dist[num] = syracuse(num)
	}
	fmt.Println(dist)
}

func syracuse(num int) int {
	count := 0
	for num != 1 {
		if num%2 == 0 {
			num /= 2
			count++
		} else {
			num = 3*num + 1
			count++
		}
	}
	return count
}
