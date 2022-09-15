package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numSix() {
	fmt.Println("Voici un nombre premier aleatoire : ")
	fmt.Print(randPrime())
}

func randPrime() int {
	var num int
	found := false
	rand.Seed(time.Now().UnixNano())
	for !found {
		num = rand.Intn(999) + 1
		if isPrime(num) {
			found = true
		}
	}

	return num
}

func isPrime(num int) bool {
	count := 0
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	return count == 0 && num != 1
}
