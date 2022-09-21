package main

import "fmt"

func main() {
	nombres := make(chan int)
	carres := make(chan int)

	go Compteurs(nombres)
	go func(out chan<- int, in <-chan int) {
		for n := range in {
			out <- n * n
		}
		defer close(out)
	}(carres, nombres)

	for x := range carres {
		fmt.Println(x)
	}
}

func Compteurs(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	defer close(out)
}


