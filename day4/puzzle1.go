package day4

import "log"

type Board struct {
	Tiles []Tile
}

type Tile struct {
	number int
	location
	selected bool
}

type location struct {
	x, y int
}

func ParseBoards(boards [][]int) []Board {
	var res []Board
	for _, board := range boards {
		var b Board
		x, y := 0, 0
		for _, boardInput := range board {
			b.Tiles = append(b.Tiles, Tile{
				number: boardInput,
				location: location{
					x: x,
					y: y,
				},
				selected: false,
			})
			x++
			if x == 5 {
				x = 0
				y++
			}
		}
		res = append(res, b)
	}
	return res
}

func CallNumbers(boards []Board, numsToCall []int) (winningBoardScore int) {
	for _, call := range numsToCall {
		// check each board and mark each number when called.
		for b, board := range boards {
			for i, tile := range board.Tiles {
				// after each check, check to see if there is match on x. if yes, call calcSum
				if tile.number == call {

					board.Tiles[i].selected = true
					//check for win condition x.
					if checkForWinX(board.Tiles, i) {
						log.Println("a win on x on board ", b)
						//calc sum and return
						return calcBoardSum(board.Tiles, call)
					}
					// check for win condition y.
					if checkForWinY(board.Tiles, i) {
						log.Println("a win on y on board", b)
						return calcBoardSum(board.Tiles, call)
					}
				}
			}
		}
		// if not, proceed through through each call.
	}

	// return 0 here if no match
	return 0
}

func checkForWinX(tiles []Tile, indexSelected int) bool {
	toCheckWinY := tiles[indexSelected].y
	yWins := 0
	for _, tile := range tiles {
		if tile.y == toCheckWinY && tile.selected {
			yWins++
		}
	}
	return yWins == 5
}

func checkForWinY(tiles []Tile, indexSelected int) bool {
	toCheckWinX := tiles[indexSelected].x
	xWins := 0
	for _, tile := range tiles {
		if tile.x == toCheckWinX && tile.selected {
			xWins++
		}
	}
	return xWins == 5
}

func calcBoardSum(tiles []Tile, winningNumber int) int {
	log.Printf("winning number %d", winningNumber)
	sum := 0
	for _, t := range tiles {
		if !t.selected {
			sum += t.number
		}
	}
	return sum * winningNumber
}
