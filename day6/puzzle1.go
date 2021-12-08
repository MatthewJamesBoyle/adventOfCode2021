package day6

import "log"

func PuzzleOne(fishys []int, days int, shouldLog bool) int {
	for d := 1; d <= days; d++ {
		for i := 0; i < len(fishys); i++ {
			fishys[i] = fishys[i] - 1
			if fishys[i] == -1 {
				fishys[i] = 6
				fishys = append(fishys, 9)
			}

		}

		if shouldLog {
			log.Printf("after %d days: %v", d, fishys)
		}
	}
	return len(fishys)

}
