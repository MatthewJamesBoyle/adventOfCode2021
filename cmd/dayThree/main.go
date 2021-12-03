package main

import (
	"github.com/matthewjamesboyle/adventOfCode2021/common"
	"github.com/matthewjamesboyle/adventOfCode2021/day3"
	"log"
)

func main() {
	const inputFile = "./day3/input.txt"

	input, err := common.ReadFileToMemory(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	res, err := day3.PuzzleTwo(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
