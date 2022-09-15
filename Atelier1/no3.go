package main

func numTree(array []int) (int, int, int) {
	var min, max, moy int
	moy = 0
	array = numTwo(array)
	min = array[0]
	max = array[len(array)-1]

	for i := 0; i < len(array); i++ {
		moy += array[i]
	}
	moy /= len(array)

	return min, max, moy
}
