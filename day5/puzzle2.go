package day5

func PuzzleTwo(input []Line) int {
	res := make(map[int]map[int]int)

	for _, line := range input {
		for _, point := range line.GetPointsOnLine(true) {
			row := res[point.X]
			if row == nil {
				res[point.X] = make(map[int]int)
				res[point.X][point.Y]++
			} else {
				res[point.X][point.Y]++
			}
		}
	}

	return CalculateResult(res)
}
