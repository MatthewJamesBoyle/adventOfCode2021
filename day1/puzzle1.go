package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	dayOneFileInput = "./day1/inputpuzzleone.txt"
)

func puzzleOne() {
	// This will be our result
	var increases = 0

	file, err := os.Open(dayOneFileInput)
	if err != nil {
		log.Fatal("failed to load file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var last int64
	for scanner.Scan() {
		if last == 0 {
			tmp, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				log.Fatal("unexpected error, one of the inputs was not parsable to an int: ", err)
			}
			last = tmp
			continue
		}
		cur, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal("unexpected error, one of the inputs was not parsable to an int: ", err)
		}
		fmt.Printf("considering last %d and current %d \n", last, cur)
		if cur > last {
			fmt.Println("considered them and increased")
			increases++
		}
		last = cur
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner broke mate", err)
	}

	fmt.Println("result:", increases)
}
