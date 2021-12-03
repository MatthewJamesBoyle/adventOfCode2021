package day3

import (
	"testing"
)

func TestPuzzleTwo(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	res, err := PuzzleTwo(input)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if res != 230 {
		t.Fatalf("wrong answer for test, got %d", res)
	}
}
