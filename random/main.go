package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	r "crypto/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	r1 := rand.Intn(10)
	fmt.Println(r1)

	r2, _ := r.Int(r.Reader, big.NewInt(10))
	fmt.Println(r2)
}
