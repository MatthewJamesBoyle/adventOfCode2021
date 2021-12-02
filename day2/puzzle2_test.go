package day2_test

import (
	"github.com/matthewjamesboyle/adventOfCode2021/day2"
	"testing"
)

func TestPuzzleTwo(t *testing.T) {
	t.Run("should return 900", func(t *testing.T) {
		input := []string{"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}

		res, err := day2.PuzzleTwo(input)
		if err != nil {
			t.Fatalf("Expected no error but got %v", err.Error())
		}

		if res != 900 {
			t.Fatalf("expected 900 but got %d", res)
		}
	})
}
