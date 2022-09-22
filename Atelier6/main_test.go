package main

import (
	"testing"
)

func BenchmarkNumOneA(b *testing.B) {

	files := getFiles()

	for i := 0; i < b.N; i++ {
		numTwoA(files)
	}
}

func BenchmarkNumOneB(b *testing.B) {

	files := getFiles()
	var c chan string = make(chan string)

	for i := 0; i < b.N; i++ {
		numTwoB(files, c)
	}
}
