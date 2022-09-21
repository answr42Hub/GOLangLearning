package main

type Rectangle struct {
	h, l float64
}

func (r *Rectangle) aire() float64 {
	return r.h * r.l
}

func (c *Rectangle) perimetre() float64 {
	return 2 * (c.h + c.l)
}
