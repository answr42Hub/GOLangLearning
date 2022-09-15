package main

import "fmt"

func numFive() {
	var num int
	fmt.Println("Combien de nombre de la suite fibonacci voulez-vous voir?")
	fmt.Scanf("%d", &num)
	fmt.Println("Voici : ")
	fmt.Print(fibonacci(num))
}

func fibonacci(num int) []int {
	array := []int{1}
	next := 0
	if num > 2 {
		array = append(array, 1)
		for i := 1; i < num-1; i++ {
			next = array[i-1] + array[i]
			array = append(array, next)
		}
	} else if num == 2 {
		array = append(array, 1)
	}

	return array
}
