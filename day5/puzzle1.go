package day5

type Line struct {
	StartOfLine *Point
	EndOfLine   *Point
}
type Point struct {
	X, Y int
}

func PuzzleOne(input []Line) int {
	res := make(map[int]map[int]int)

	for _, line := range input {
		for _, point := range line.GetPointsOnLine(false) {
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

func CalculateResult(resMap map[int]map[int]int) int {
	counter := 0
	for _, row := range resMap {
		for _, column := range row {
			if column >= 2 {
				counter++
			}
		}
	}
	return counter
}

func (l *Line) GetPointsOnLine(canDiag bool) []*Point {
	points := make([]*Point, 0)

	if l.StartOfLine.X != l.EndOfLine.X && l.StartOfLine.Y != l.EndOfLine.Y && !canDiag {
		return points
	}

	curr := l.StartOfLine
	for curr != nil {
		points = append(points, curr)
		curr = curr.NextPoint(l.EndOfLine)
	}
	return points
}

func (curr *Point) NextPoint(targetPoint *Point) *Point {
	if curr.X == targetPoint.X && curr.Y == targetPoint.Y {
		return nil
	}
	return &Point{
		X: curr.X + getAdditionValue(curr.X, targetPoint.X),
		Y: curr.Y + getAdditionValue(curr.Y, targetPoint.Y),
	}
}

func getAdditionValue(p1, p2 int) int {
	switch {
	case p1 == p2:
		return 0
	case p1 < p2:
		return 1
	default:
		return -1
	}
}
