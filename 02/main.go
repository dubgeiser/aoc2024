package main

import (
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"math"
	"slices"
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
	for _, r := range s.reports {
		valid := false
		for i := 0; i < len(r); i++ {
			rr := slices.Concat(r[:i], r[i+1:])
			if isSafe(rr) {
				valid = true
			}
		}
		if valid {
			total++
		}
	}
	return total
}

func distance(a, b int) int {
	return int(math.Abs(float64(b - a)))
}

func isSafe(r []int) bool {
	for i := 1; i < len(r); i++ {
		if distance(r[i-1], r[i]) < 1 || distance(r[i-1], r[i]) > 3 {
			return false
		}
	}
	rs := slices.Sorted(slices.Values(r))
	rr := slices.Clone(rs)
	slices.Reverse(rr)
	return slices.Equal(r, rs) || slices.Equal(r, rr)
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
