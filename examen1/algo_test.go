package main

import (
	"math/rand"
	"testing"
	"time"
)

var tests = []paires{
	{50, 28},
	{100, 45},
	{1000, 496},
	{0, 0},
	{1, 0},
}

func TestAlgo(t *testing.T) {
	for _, test := range tests {
		got := Algo(test.n)
		if got != test.result {
			t.Errorf("Algo(%d) = %d", test.n, got)
		}
	}
}

func BenchmarkAlgo(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		Algo(rand.Intn(100001))
	}
}

type paires struct {
	n      int
	result int
}

/*
   goos: linux
   goarch: amd64
   pkg: examen1
   cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
   BenchmarkAlgo
   BenchmarkAlgo-8   	     9107194	           112.3 ns/op
   PASS
*/
