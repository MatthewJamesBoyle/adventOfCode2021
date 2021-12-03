package day3

import (
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	t.Run("most common and least common are correct", func(t *testing.T) {
		tests := []struct {
			input           string
			wantMostCommon  string
			wantLeastCommon string
		}{
			{input: "00100", wantMostCommon: "0", wantLeastCommon: "1"},
			{input: "11110", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "10110", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "10111", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "00111", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "11100", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "10000", wantMostCommon: "0", wantLeastCommon: "1"},
			{input: "11001", wantMostCommon: "1", wantLeastCommon: "0"},
			{input: "00010", wantMostCommon: "0", wantLeastCommon: "1"},
			{input: "01010", wantMostCommon: "0", wantLeastCommon: "1"},
		}

		for _, v := range tests {
			gotmost, gotleast := findMostCommonAndLeastCommon(v.input)
			if gotmost != v.wantMostCommon {
				t.Fatalf("expected %s but got %s for most common occ", v.wantMostCommon, gotmost)
			}

			if gotleast != v.wantLeastCommon {
				t.Fatalf("expected %s but got %s for least common occ", v.wantLeastCommon, gotleast)
			}
		}
	})

	t.Run("output correct given input", func(t *testing.T) {
		inp := []string{
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

		res, err := PuzzleOne(inp)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if res != 198 {
			t.Fatalf("expected res to be 198, got %d", res)
		}
	})

}
