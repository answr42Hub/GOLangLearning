package main

func numTwob(array []int) []int {

	v1 := array[:len(array)-1]

	for i := range v1 {
		v2 := array[i+1:]
		for j := range v2 {
			if v1[i] > v2[j] {
				v1[i], v2[j] = v2[j], v1[i]
			}
		}
	}
	return array
}
