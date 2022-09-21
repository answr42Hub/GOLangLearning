package main

import "fmt"

func main() {
	multiForme := MultiForm{
		formes: []Forme{
			&Cercle{0.0, 0.0, 5.0},
			&Rectangle{10.0, 10.0},
		},
	}

	fmt.Println(multiForme.aire())
	fmt.Println(multiForme.perimetre())
}
