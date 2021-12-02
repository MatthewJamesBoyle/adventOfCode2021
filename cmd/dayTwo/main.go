package main

import (
	"bufio"
	"github.com/matthewjamesboyle/adventOfCode2021/day2"
	"log"
	"os"
)

func main() {
	const inputFile = "./day2/input.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	res, err := day2.PuzzleOne(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
