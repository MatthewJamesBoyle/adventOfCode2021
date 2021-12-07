package main

import (
	"fmt"
	"github.com/matthewjamesboyle/adventOfCode2021/day5"
	"log"
)

func main() {
	const inputFile = "./day5/input.txt"

	grid, err := day5.ReadFileToLines(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(day5.PuzzleTwo(grid))

}
