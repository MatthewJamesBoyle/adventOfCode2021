package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type window struct {
	w []int64
}

const windowSize = 3

func (w *window) populate(val int64) {
	if len(w.w) < windowSize {
		w.w = append(w.w, val)
		return
	}
	if len(w.w) == windowSize {
		w.w[0] = w.w[1]
		w.w[1] = w.w[2]
		w.w[2] = val
	}
}

func (w *window) isFull() bool {
	return len(w.w) == windowSize
}

func (w *window) sum() int64 {
	res := int64(0)
	for _, v := range w.w {
		res += v
	}
	return res
}

func PuzzleTwo(fileLocation string) {
	// This will be our result
	var increases = 0

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal("failed to load file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var w window
	lastWindowVal := int64(0)
	for scanner.Scan() {
		tmp, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal("unexpected error, one of the inputs was not parsable to an int: ", err)
		}
		w.populate(tmp)

		if !w.isFull() {
			continue
		}

		// this is for the case where we have no previous window
		if lastWindowVal == 0 {
			lastWindowVal = w.sum()
			continue
		}

		if w.sum() > lastWindowVal {
			log.Printf("%d increased", w.sum())
			increases++
		}
		if w.sum() < lastWindowVal {
			log.Printf("%d decreased", w.sum())
		}

		if w.sum() == lastWindowVal {
			log.Printf("%d no change", w.sum())
		}
		lastWindowVal = w.sum()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner broke mate", err)
	}

	fmt.Println("result:", increases)
}
