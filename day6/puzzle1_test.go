package day6

import (
	"log"
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	t.Run("test with example input", func(t *testing.T) {
		fish := []int{3, 4, 3, 1, 2}

		log.Println(PuzzleOne(fish, 80, false))
	})
}
