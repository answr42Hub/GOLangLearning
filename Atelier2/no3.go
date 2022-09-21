package main

func numTree(sum *int, args ...int) {
	for _, arg := range args {
		*sum += arg
	}
}
