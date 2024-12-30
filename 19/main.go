package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func numCombinations(towels []string, design string, cache map[string]int) int {
	if ccount, ok := cache[design]; ok {
		return ccount
	}
	count := 0
	if len(design) == 0 {
		count = 1
	}
	for _, towel := range towels {
		reduced, found := strings.CutPrefix(design, towel)
		if found {
			count += numCombinations(towels, reduced, cache)
		}
	}
	cache[design] = count
	return count
}



func main() {
	towels := []string{}
	designs := []string{}
	input.TwoParts(
		func(s *bufio.Scanner) { towels = slices.Concat(towels, strings.Split(s.Text(), ", ")) },
		func(s *bufio.Scanner) { designs = append(designs, s.Text()) })
	p1 := 0
	p2 := 0
	cache := map[string]int{}
	for _, design := range designs {
		result := numCombinations(towels, design, cache)
		if result > 0 {
			p1++
			p2 += result
		}
	}
	fmt.Println(p1, p2)
}
