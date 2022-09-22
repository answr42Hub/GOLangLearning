package main

import (
	m "raph/math"
	"testing"
)

type paires struct {
	param   []int
	reponse float64
}

// Fichier math_test.go
// go test to run the test of sum
var valeursTestSum = []paires{
	{[]int{0, 0}, 0},
	{[]int{10, 55}, 65},
	{[]int{4, 3}, 7},
	{[]int{4, 3, 2}, 9},
	{[]int{}, 0},
}

// Fichier math_test.go
// go test to run the test of moy
var valeursTestMoy = []paires{
	{[]int{0, 0, 0, 0}, 0},
	{[]int{10, 55}, 32.5},
	{[]int{4, 3}, 3.5},
	{[]int{4, 3, 2}, 3},
	{[]int{}, 0},
}

// Fichier math_test.go
// go test to run the test of stddev
var valeursTestStddev = []paires{
	{[]int{0, 0, 0, 0}, 0},
	{[]int{10, 55}, 22.5},
	{[]int{4, 3}, 0.5},
	{[]int{4, 3, 2}, 0.816496580927726},
	{[]int{}, 0},
}

// Fichier math_test.go
// go test to run the test of min
var valeursTestMin = []paires{
	{[]int{0, 0, 0, 0}, 0},
	{[]int{10, 55}, 10},
	{[]int{4, 3}, 3},
	{[]int{4, 3, 2}, 2},
	{[]int{}, 0},
}

// Fichier math_test.go
// go test to run the test of max
var valeursTestMax = []paires{
	{[]int{0, 0, 0, 0}, 0},
	{[]int{10, 55}, 55},
	{[]int{4, 3}, 4},
	{[]int{4, 3, 2}, 4},
	{[]int{}, 0},
}

// Fichier math_test.go
// go test to run the test of med
var valeursTestMed = []paires{
	{[]int{0, 0, 0, 0}, 0},
	{[]int{10, 55}, 32.5},
	{[]int{4, 3}, 3.5},
	{[]int{4, 3, 2}, 3},
	{[]int{}, 0},
}

func TestSum(t *testing.T) {
	for _, paire := range valeursTestSum {
		f := m.Sum(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}

func TestMoy(t *testing.T) {
	for _, paire := range valeursTestMoy {
		f := m.Moy(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}

func TestStddev(t *testing.T) {
	for _, paire := range valeursTestStddev {
		f := m.Stddev(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}

func TestMin(t *testing.T) {
	for _, paire := range valeursTestMin {
		f := m.Min(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}

func TestMax(t *testing.T) {
	for _, paire := range valeursTestMax {
		f := m.Max(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}

func TestMed(t *testing.T) {
	for _, paire := range valeursTestMed {
		f := m.Med(paire.param)
		if f != paire.reponse {
			t.Errorf("Attendu %f, Obtenu %f", paire.reponse, f)
		}
	}
}
