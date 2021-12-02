package day2_test

import (
	"github.com/matthewjamesboyle/adventOfCode2021/day2"
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	t.Run("should return 150", func(t *testing.T) {
		input := []string{"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}

		res, err := day2.PuzzleOne(input)
		if err != nil {
			t.Fatalf("Expected no error but got %v", err.Error())
		}

		if res != 150 {
			t.Fatalf("expected 150 but got %d", res)
		}
	})
}
