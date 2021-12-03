package day3

import (
	"fmt"
	"golang.org/x/sync/errgroup"

	"strconv"
	"strings"
)

type rec struct {
	original string
	rec      string
}

const (
	oxy = iota
	co2 = iota
)

func PuzzleTwo(input []string) (int64, error) {
	var recSlice []rec
	for _, v := range input {
		recSlice = append(recSlice, rec{
			original: v,
			rec:      v,
		})
	}
	var eg errgroup.Group

	var oxygenRes, coRes int64
	eg.Go(func() error {
		ox, err := calculate(recSlice, oxy)
		if err != nil {
			return fmt.Errorf("error oxygen calc: %w", err)
		}
		oxygenRes = ox

		return nil
	})

	eg.Go(func() error {
		co, err := calculate(recSlice, co2)
		if err != nil {
			return fmt.Errorf("error co2 calc: %w", err)
		}
		coRes = co
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 0, err
	}

	return oxygenRes * coRes, nil
}

func calculate(input []rec, which int) (int64, error) {
	if len(input) == 1 {
		res, err := strconv.ParseInt(input[0].original, 2, 64)
		if err != nil {
			return 0, err
		}
		return res, nil
	}

	zeroCount := 0
	oneCount := 0
	for _, v := range input {
		bits := strings.Split(v.rec, "")
		firstChar := bits[0]
		if firstChar == "0" {
			zeroCount++
		}
		if firstChar == "1" {
			oneCount++
		}

	}

	var res []rec
	if which == oxy {
		res = filterForOxygen(input, zeroCount, oneCount)
	} else {
		res = filterForCO2(input, zeroCount, oneCount)
	}

	return calculate(res, which)
}

func filterForOxygen(input []rec, zeroCount, oneCount int) []rec {
	var res []rec
	if zeroCount > oneCount {
		res = filter(input, "0")
	}
	if oneCount > zeroCount || zeroCount == oneCount {
		res = filter(input, "1")
	}
	return res
}

func filterForCO2(input []rec, zeroCount, oneCount int) []rec {
	var res []rec
	if zeroCount < oneCount || zeroCount == oneCount {
		res = filter(input, "0")
	}
	if oneCount < zeroCount {
		res = filter(input, "1")
	}
	return res
}

func filter(toFilter []rec, bit string) []rec {
	var res []rec
	for _, v := range toFilter {
		firstChar := v.rec[0:1]
		if firstChar == bit {
			res = append(res, rec{
				original: v.original,
				rec:      v.rec[1:],
			})
		}
	}
	return res
}
