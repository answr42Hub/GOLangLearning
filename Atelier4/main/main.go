package main

import (
	"flag"
	"fmt"
	m "raph/math"
	"strconv"
	"strings"
)

func main() {
	var mathFunc = flag.String("c", "", "-c=[] : sum for sum, ")
	var numbers = flag.String("n", "", "enter numbers")
	flag.Parse()

	nums := strings.Split(*numbers, ",")

	ints := make([]int, len(nums))

	for i, s := range nums {
		ints[i], _ = strconv.Atoi(s)
	}

	switch *mathFunc {
	case "sum":
		fmt.Println(m.Sum(ints))
	case "moy":
		fmt.Println(m.Moy(ints))
	case "min":
		fmt.Println(m.Min(ints))
	case "max":
		fmt.Println(m.Max(ints))
	case "med":
		fmt.Println(m.Med(ints))
	case "std":
		fmt.Println(m.Stddev(ints))
	}
}
