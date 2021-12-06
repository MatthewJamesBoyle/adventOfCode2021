package day4

func FindLastWinner(boards []Board, numsToCall []int) (winningBoardScore int) {
	lastCalled := 0
	var lastWinningBoard Board
	for _, call := range numsToCall {
		// check each board and mark each number when called.
		for b, board := range boards {
			if board.completed {
				continue
			}
			for i, tile := range board.Tiles {
				// after each check, check to see if there is match on x. if yes, call calcSum
				if tile.number == call {
					board.Tiles[i].selected = true
					//check for win condition x.
					if checkForWinX(board.Tiles, i) {
						lastWinningBoard = board
						lastCalled = call
						boards[b].completed = true
					}
					// check for win condition y.
					if checkForWinY(board.Tiles, i) {
						lastWinningBoard = board
						lastCalled = call
						boards[b].completed = true
					}
				}
			}
		}
		// if not, proceed through through each call.
	}
	return calcBoardSum(lastWinningBoard.Tiles, lastCalled)
}
