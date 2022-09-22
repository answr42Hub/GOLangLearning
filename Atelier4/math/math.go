package _math

import (
	"math"
	"sort"
)

func Sum(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum)
}

func Moy(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	return float64(Sum(nums)) / float64(len(nums))
}

func Stddev(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	moy := Moy(nums)
	sum := 0.0
	for _, n := range nums {
		sum += math.Pow(float64(n)-moy, 2)
	}
	return math.Sqrt(sum / float64(len(nums)))
}

func Min(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	sort.Ints(nums)
	return float64(nums[0])
}

func Max(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	sort.Ints(nums)
	return float64(nums[len(nums)-1])
}

func Med(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	sort.Ints(nums)
	slen := len(nums)
	if slen%2 == 0 {
		return (float64(nums[slen/2-1]) + float64(nums[slen/2])) / 2
	}
	return float64(nums[slen/2])
}
