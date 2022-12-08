package main

import "fmt"

func main() {
	fmt.Println(Algo(50))
	fmt.Println(Algo(100))
	fmt.Println(Algo(1000))
	fmt.Println(Algo(0))
	fmt.Println(Algo(1))
}

func Algo(n int) int {
	var sum int
	for i := 1; i*i < n; i++ {
		sum += i
	}
	return sum
}
