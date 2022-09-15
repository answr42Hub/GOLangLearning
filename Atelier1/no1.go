package main

import "fmt"

func numOne() {
	fmt.Println("Entrez un nombre réel :")
	var num float32
	fmt.Scanf("%f", &num)

	fmt.Printf("There you go %f\n", num*2)

}
