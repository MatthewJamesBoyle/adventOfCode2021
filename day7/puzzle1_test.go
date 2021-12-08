package day7

import (
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	t.Run("correct output given test input", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		fuelUsed := PuzzleOne(input)
		if fuelUsed != 37 {
			t.Fatalf("expected 37 but got %d", fuelUsed)
		}
	})
}
