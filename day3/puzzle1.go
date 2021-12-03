package day3

import (
	"strconv"
	"strings"
)

func PuzzleOne(input []string) (int64, error) {
	occurrenceCounter := make(map[int]string, 0)
	for _, v := range input {
		bits := strings.Split(v, "")
		for i, b := range bits {
			curVal := occurrenceCounter[i]
			occurrenceCounter[i] = curVal + b
		}
	}

	var gamma, epsilon string
	for i := 0; i < len(occurrenceCounter); i++ {
		most, least := findMostCommonAndLeastCommon(occurrenceCounter[i])
		gamma += most
		epsilon += least
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return 0, err
	}

	return gammaDec * epsilonDec, nil
}

func findMostCommonAndLeastCommon(binaryInput string) (mostcommon string, leastcommon string) {
	zeroCounter := 0
	oneCounter := 0
	bits := strings.Split(binaryInput, "")

	for _, v := range bits {
		if v == "0" {
			zeroCounter++
		}
		if v == "1" {
			oneCounter++
		}
	}

	if zeroCounter > oneCounter {
		return "0", "1"
	}
	if oneCounter > zeroCounter {
		return "1", "0"
	}

	if zeroCounter == oneCounter {
		// not clear from the puzzle what to do here.
	}
	return "0", "0"
}
