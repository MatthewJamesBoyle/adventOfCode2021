package day7

import (
	"math"
	"sort"
)

func PuzzleOne(input []int) int {
	sort.Ints(input)
	var mid int
	if len(input)%2 == 0 {
		mid = input[len(input)/2]
	} else {
		mid = input[len(input)/2] + input[(len(input)-1)]/2
	}

	fuel := 0
	for _, v := range input {
		diff := int(math.Abs(float64(v) - float64(mid)))
		fuel += diff
	}
	return fuel
}
