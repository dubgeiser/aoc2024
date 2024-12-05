package main

import (
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"slices"
	"strings"
)

const MODE_RULES = 1
const MODE_UPDATES = 2

type Solution struct {
	parseMode int
	before    map[int][]int
	updates   [][]int
}

func (s *Solution) ProcessLine(i int, line string) {
	if line == "" {
		s.parseMode = MODE_UPDATES
		return
	}
	if s.parseMode == MODE_RULES {
		pages := []int(slice.Map(slice.Int, strings.Split(line, "|")))
		if len(pages) != 2 {
			fmt.Println("Rule", pages, "has", len(pages), "pages")
		}
		// map page to a list of pages that must come before it
		s.before[pages[1]] = append(s.before[pages[1]], pages[0])
	} else {
		pages := []int(slice.Map(slice.Int, strings.Split(line, ",")))
		if len(pages)%2 == 0 {
			fmt.Println("Update", pages, "has", len(pages), "pages")
		}
		s.updates = append(s.updates, pages)
	}
}

func (s *Solution) isValid(u []int) bool {
	valid := true
	for i, p1 := range u {
		for j, p2 := range u {
			// if we compare 2 pages where `p1` comes before `p2`
			// and
			// `p2` is in the list of pages that come _before_ `p1`
			// we have an invalid page update.
			if i < j && slices.Contains(s.before[p1], p2) {
				valid = false
			}
		}
	}
	return valid
}

func (s *Solution) Part1() any {
	result := 0
	for _, u := range s.updates {
		if s.isValid(u) {
			result += u[len(u)/2]
		}
	}
	return result
}

func (s *Solution) Part2() any {
	result := 0
	for _, u := range s.updates {
		if !s.isValid(u) {
			us := slices.SortedFunc(slices.Values(u), func(p1, p2 int) int {
				if slices.Contains(s.before[p1], p2) {
					return -1
				}
				if slices.Contains(s.before[p2], p1) {
					return 1
				}
				return 0
			})
			result += us[len(us)/2]
		}
	}
	return result
}

func newSolution() *Solution {
	s := &Solution{}
	s.parseMode = MODE_RULES
	s.before = make(map[int][]int)
	return s
}

func main() {
	s := newSolution()
	s.parseMode = MODE_RULES
	n, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Read", n, "lines")
	fmt.Println("Part1:", s.Part1())
	fmt.Println("Part2:", s.Part2())
}
