package main

import (
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"math"
	"strings"
)

type Solution struct {
	reports [][]int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.reports = append(s.reports, slice.Map(slice.Int, strings.Split(line, " ")))
}

func (s *Solution) Part1() int {
	total := 0
	for _, report := range s.reports {
		if isSafe(report) {
			total++
		}
	}
	return total
}

func (s *Solution) Part2() int {
	total := 0
	var rr []int
	for _, report := range s.reports {
		ok := false
		for i := 0; i < len(report); i++ {
			rr = tolerate(report, i)
			if isSafe(rr) {
				ok = true
			}
		}
		if ok {
			total++
		}
	}
	return total
}

func tolerate(report []int, iLevel int) []int {
	var rr []int
	for i, level := range report {
		if i == iLevel {
			continue
		}
		rr = append(rr, level)
	}
	return rr
}

const ASCENDING = 1
const DESCENDING = 2
const EQUAL = 3

func determineOrder(prev, next int) int {
	if prev < next {
		return ASCENDING
	} else if prev > next {
		return DESCENDING
	} else {
		return EQUAL
	}
}

func distance(a, b int) int {
	return int(math.Abs(float64(b - a)))
}

func isSafe(report []int) bool {
	order := determineOrder(report[0], report[1])
	if order == EQUAL {
		return false
	}
	prevLevel := report[0]
	for i := 1; i < len(report); i++ {
		currLevel := report[i]
		currOrder := determineOrder(prevLevel, currLevel)
		if currOrder != order {
			return false
		}
		if distance(prevLevel, currLevel) < 1 || distance(prevLevel, currLevel) > 3 {
			return false
		}
		prevLevel = currLevel
	}
	return true
}

func main() {
	s := &Solution{}
	_, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Part1:", s.Part1())
	fmt.Println("Part2:", s.Part2())
}
