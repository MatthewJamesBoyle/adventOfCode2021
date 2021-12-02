package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	amount    int64
}

const (
	directionForward = "forward"
	directionDown    = "down"
	directionUp      = "up"
)

func PuzzleOne(input []string) (int64, error) {
	depth := int64(0)
	horizontalPosition := int64(0)

	ins, err := buildInstructionSet(input)
	if err != nil {
		return 0, fmt.Errorf("failed to build instructions: %w", err)
	}
	for _, v := range ins {
		switch v.direction {
		// forward 5 adds 5 to your horizontal position, a total of 5.
		case directionForward:
			horizontalPosition += v.amount
			//down 5 adds 5 to your depth, resulting in a value of 5.
		case directionDown:
			depth += v.amount
			//up 3 decreases your depth by 3, resulting in a value of 2
		case directionUp:
			depth -= v.amount
		default:
			return 0, fmt.Errorf("unexpected instruction %s", v.direction)
		}
	}
	return horizontalPosition * depth, nil
}

func buildInstructionSet(input []string) ([]instruction, error) {
	var ins = make([]instruction, 0,len(input))
	for _, v := range input {
		res := strings.Split(v, " ")
		if len(res) != 2 {
			return nil, errors.New("failed to split the instruction")
		}
		amt, err := strconv.ParseInt(res[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int: %w", err)
		}
		ins = append(ins, instruction{direction: res[0], amount: amt})
	}
	return ins, nil
}
