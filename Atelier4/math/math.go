package _math

import (
	"math"
	"sort"
)

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Moy(nums []int) float64 {
	return float64(Sum(nums)) / float64(len(nums))
}

func Stddev(nums []int) float64 {

	moy := Moy(nums)
	sum := 0.0
	for _, n := range nums {
		sum += math.Pow(float64(n)-moy, 2)
	}
	return math.Sqrt(sum / float64(len(nums)))
}

func Min(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func Max(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func Med(nums []int) float64 {
	sort.Ints(nums)
	slen := len(nums)
	if slen%2 == 0 {
		return (float64(nums[slen/2-1]) + float64(nums[slen/2])) / 2
	}
	return float64(nums[slen/2])
}
