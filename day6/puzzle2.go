package day6

import "log"

func PuzzleTwo(fishys []int, days int, shouldLog bool) int {
	for d := 1; d <= days; d++ {
		for i := 0; i < len(fishys); i++ {
			fishys[i] = fishys[i] - 1
			if fishys[i] == -1 {
				fishys[i] = 6
				fishys = append(fishys, 9)
			}
		}
		if shouldLog {
			log.Println("finished day:", d)
		}
	}
	return len(fishys)
}
