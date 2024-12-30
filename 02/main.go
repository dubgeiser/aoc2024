package main

import (
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"
)

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
	reports := [][]int{}
	input.Lines(func(s *bufio.Scanner) {
		reports = append(reports, slice.Map(slice.Int, strings.Split(s.Text(), " ")))
	})
	p1 := 0
	for _, report := range reports {
		if isSafe(report) {
			p1++
		}
	}

	p2 := 0
	for _, r := range reports {
		valid := false
		for i := 0; i < len(r); i++ {
			rr := slices.Concat(r[:i], r[i+1:])
			if isSafe(rr) {
				valid = true
			}
		}
		if valid {
			p2++
		}
	}

	fmt.Println(p1, p2)
}
