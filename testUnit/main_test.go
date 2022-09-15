package testUnit

import "testing"

type paires struct {
	param, reponse int
}

// Fichier main_test.go
// go test to run the test
var valeurTests = []paires{
	{0, 0},
	{10, 55},
	{4, 3},
}

func TestFibR(t *testing.T) {

	for _, paire := range valeurTests {
		f := FibR(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %d, Obtenu %d", paire.reponse, f)
		}
	}
}

func TestFibT(t *testing.T) {

	for _, paire := range valeurTests {
		f := FibT(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %d, Obtenu %d", paire.reponse, f)
		}
	}
}

func BenchmarkFibR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibR(10)
	}
}

func BenchmarkFibT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibT(10)
	}
}

func benchmarkFibR(valeur int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibR(valeur)
	}
}

func BenchmarkFibR1(b *testing.B) {
	benchmarkFibR(1, b)
}

func BenchmarkFibR2(b *testing.B) {
	benchmarkFibR(2, b)
}

func BenchmarkFibR10(b *testing.B) {
	benchmarkFibR(10, b)
}

func BenchmarkFibR50(b *testing.B) {
	benchmarkFibR(50, b)
}

func benchmarkFibT(valeur int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibT(valeur)
	}
}

func BenchmarkFibT1(b *testing.B) {
	benchmarkFibT(1, b)
}

func BenchmarkFibT2(b *testing.B) {
	benchmarkFibT(2, b)
}

func BenchmarkFibT10(b *testing.B) {
	benchmarkFibT(10, b)
}

func BenchmarkFibT50(b *testing.B) {
	benchmarkFibT(50, b)
}
