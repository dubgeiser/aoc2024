package main

import (
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func isValid(u []int, before map[int][]int) bool {
	valid := true
	for i, p1 := range u {
		for j, p2 := range u {
			// if we compare 2 pages where `p1` comes before `p2`
			// and
			// `p2` is in the list of pages that come _before_ `p1`
			// we have an invalid page update.
			if i < j && slices.Contains(before[p1], p2) {
				valid = false
			}
		}
	}
	return valid
}

func main() {
	before := make(map[int][]int)
	updates := [][]int{}
	input.TwoParts(func(s *bufio.Scanner) {
		pages := [2]int(slice.Map(slice.Int, strings.Split(s.Text(), "|")))
		// map page to a list of pages that must come before it
		before[pages[1]] = append(before[pages[1]], pages[0])
	}, func(s *bufio.Scanner) {
		pages := []int(slice.Map(slice.Int, strings.Split(s.Text(), ",")))
		if len(pages)%2 == 0 {
			fmt.Println("Update", pages, "has", len(pages), "pages")
		}
		updates = append(updates, pages)
	})

	p1 := 0
	for _, u := range updates {
		if isValid(u, before) {
			p1 += u[len(u)/2]
		}
	}

	p2 := 0
	for _, u := range updates {
		if !isValid(u, before) {
			us := slices.SortedFunc(slices.Values(u), func(p1, p2 int) int {
				if slices.Contains(before[p1], p2) {
					return -1
				}
				if slices.Contains(before[p2], p1) {
					return 1
				}
				return 0
			})
			p2 += us[len(us)/2]
		}
	}

	fmt.Println(p1, p2)
}
