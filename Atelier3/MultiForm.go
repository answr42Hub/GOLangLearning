package main

type MultiForm struct {
	formes []Forme
}

func (m *MultiForm) aire() float64 {
	var a float64
	for _, f := range m.formes {
		a += f.aire()
	}
	return a
}

func (m *MultiForm) perimetre() float64 {
	var p float64
	for _, f := range m.formes {
		p += f.perimetre()
	}
	return p
}
