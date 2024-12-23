package main

import (
	"aoc/lib/file"
	"fmt"
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
	fmt.Println()
	t, d := file.ReadTwoParts("./input")
	towels := strings.Split(t, ", ")
	designs := strings.Split(d, "\n")
	count := 0
	cache := map[string]int{}
	p2 := 0
	for _, design := range designs {
		result := numCombinations(towels, design, cache)
		if result > 0 {
			count++
			p2 += result
		}
	}
	fmt.Println(count, p2)
}
