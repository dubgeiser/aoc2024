package main

import (
	"aoc/lib/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct {
	reports [][]int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.reports = append(s.reports, toInts(line))
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
	return total
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
	for i := 1; i<len(report); i++ {
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

func toInts(line string) []int {
	l := []int{}
	for _, s := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(string(s))
		l = append(l, n)
	}
	return l
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
