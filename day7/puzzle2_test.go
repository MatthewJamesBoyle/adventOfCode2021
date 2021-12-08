package day7

import (
	"testing"
)

func TestPuzzleTwo(t *testing.T) {
	t.Run("correct output given test input", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		fuelUsed := PuzzleTwo(input)
		if fuelUsed != 168 {
			t.Fatalf("expected 168 but got %d", fuelUsed)
		}
	})
}
