package main

import "fmt"

func numFour() {
	fmt.Println("Entrez un entier :")
	var num int
	fmt.Scanf("%d", &num)
	syracuse(num)
}

func syracuse(num int) {
	fmt.Print(num)
	for num != 1 {
		fmt.Print(", ")
		if num%2 == 0 {
			num /= 2
			fmt.Print(num)
		} else {
			num = 3*num + 1
			fmt.Print(num)
		}
	}
}
