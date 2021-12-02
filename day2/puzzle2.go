package day2

import "fmt"

func PuzzleTwo(input []string) (int64, error) {
	depth := int64(0)
	horizontalPosition := int64(0)
	aim := int64(0)

	ins, err := buildInstructionSet(input)
	if err != nil {
		return 0, err
	}

	for _, v := range ins {
		switch v.direction {
		case directionForward:
			//forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
			horizontalPosition += v.amount
			if aim > 0 {
				//forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
				depth += v.amount * aim
			}
		case directionDown:
			//down 8 adds 8 to your aim, resulting in a value of 10.
			aim += v.amount
		case directionUp:
			//up 3 decreases your aim by 3, resulting in a value of 2.
			aim -= v.amount
		default:
			return 0, fmt.Errorf("unexpected instruction %s", v.direction)
		}
	}
	return horizontalPosition * depth, nil

}
