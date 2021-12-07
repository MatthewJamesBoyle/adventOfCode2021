package day5

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToLines(inputFile string) ([]Line, error) {
	var res []Line
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		points := strings.Split(line, "->")
		if len(points) != 2 {
			return nil, errors.New("didnt split on -> correctly")
		}
		leftP := strings.Split(points[0], ",")
		rightP := strings.Split(points[1], ",")
		if len(leftP) != 2 {
			return nil, errors.New("length of left was not 2")
		}
		if len(rightP) != 2 {
			return nil, errors.New("length of right was not 2")
		}

		leftX, err := strconv.ParseInt(strings.Trim(leftP[0], " "), 10, 64)
		if err != nil {
			return nil, errors.New("left x")
		}
		leftY, err := strconv.ParseInt(strings.Trim(leftP[1], " "), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("left y:%w", err)
		}

		rightX, err := strconv.ParseInt(strings.Trim(rightP[0], " "), 10, 64)
		if err != nil {
			return nil, errors.New("right x")
		}
		rightY, err := strconv.ParseInt(strings.Trim(rightP[1], " "), 10, 64)
		if err != nil {
			return nil, errors.New("right y")
		}

		res = append(res, Line{
			StartOfLine: &Point{
				X: int(leftX),
				Y: int(leftY),
			},
			EndOfLine: &Point{
				X: int(rightX),
				Y: int(rightY),
			},
		})
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("ReadFileToMemory: %w", err)
	}
	return res, nil
}
