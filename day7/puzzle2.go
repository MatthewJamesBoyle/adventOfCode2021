package day7

import (
	"math"
)

func PuzzleTwo(input []int) int {
	// find the average
	sum := 0
	for _, v := range input {
		sum += v
	}
	avg := int(math.Floor(float64(sum) / float64(len(input))))

	fuel := 0
	for _, v := range input {
		//triangle numbers
		diff := int(math.Abs(float64(v) - float64(avg)))
		fuel += diff * (diff + 1) / 2
	}
	return fuel
}
