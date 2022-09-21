package main

import "math"

type Cercle struct {
	x, y, r float64
}

func (c *Cercle) aire() float64 {
	return math.Pi * c.r * c.r
}

func (c *Cercle) perimetre() float64 {
	return 2 * math.Pi * c.r
}
