package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileToMemory(inputFile string) ([]string, error) {
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
		return nil, fmt.Errorf("ReadFileToMemory: %w")
	}
	return input, nil
}
