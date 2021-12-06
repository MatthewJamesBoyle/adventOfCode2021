package day4

import (
	"testing"
)

func TestParseBoard(t *testing.T) {
	board1 := []int{
		22, 13, 1, 7, 11, 0,
		8, 2, 23, 4, 24,
		21, 9, 14, 16, 7,
		6, 10, 3, 18, 5,
		1, 12, 20, 15, 19,
	}

	board2 := []int{
		3, 15, 0, 2, 22,
		9, 18, 13, 17, 5,
		19, 8, 7, 25, 23,
		20, 11, 10, 24, 4,
		14, 21, 16, 12, 6,
	}

	board3 := []int{
		14, 21, 17, 24, 4,
		10, 16, 15, 9, 19,
		18, 8, 23, 26, 20,
		22, 11, 13, 6, 5,
		2, 0, 12, 3, 7,
	}

	res := ParseBoards([][]int{board1, board2, board3})

	if len(res) != 3 {
		t.Fatalf("expected to get %d boards but got %d", 3, len(res))
	}

	// random sample
	toSample1 := res[0].Tiles[3]
	toSample2 := res[1].Tiles[9]
	toSample3 := res[2].Tiles[11]

	if toSample1.selected != false {
		t.Fail()
	}
	if toSample1.number != 7 {
		t.Fail()
	}
	if toSample1.x != 3 {
		t.Fail()
	}
	if toSample1.y != 0 {
		t.Fail()
	}

	if toSample2.selected != false {
		t.Fail()
	}
	if toSample2.number != 5 {
		t.Fatalf("expected 5 but got %d", toSample2.number)
	}
	if toSample2.x != 4 {
		t.Fatalf("expected x to be 4 but got %d", toSample2.x)

	}
	if toSample2.y != 1 {
		t.Fatalf("expected y to be 1 but got %d", toSample2.y)
	}

	if toSample3.selected != false {
		t.Fail()
	}
	if toSample3.number != 8 {
		t.Fatalf("expected 8 but got %d", toSample3.number)
	}
	if toSample3.x != 1 {
		t.Fatalf("expected x to be 3 but got %d", toSample3.x)

	}
	if toSample3.y != 2 {
		t.Fatalf("expected y to be 2 but got %d", toSample3.y)
	}
}

func TestCallNumbers(t *testing.T) {
	board1 := []int{
		22, 13, 1, 7, 11, 0,
		8, 2, 23, 4, 24,
		21, 9, 14, 16, 7,
		6, 10, 3, 18, 5,
		1, 12, 20, 15, 19,
	}

	board2 := []int{
		3, 15, 0, 2, 22,
		9, 18, 13, 17, 5,
		19, 8, 7, 25, 23,
		20, 11, 10, 24, 4,
		14, 21, 16, 12, 6,
	}

	board3 := []int{
		14, 21, 17, 24, 4,
		10, 16, 15, 9, 19,
		18, 8, 23, 26, 20,
		22, 11, 13, 6, 5,
		2, 0, 12, 3, 7,
	}

	boards := ParseBoards([][]int{board1, board2, board3})
	toCall := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	sum := CallNumbers(boards, toCall)

	if sum != 4512 {
		t.Fatalf("expected 4512 but got %d", sum)
	}

}
