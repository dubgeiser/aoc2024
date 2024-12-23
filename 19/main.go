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
	for _, design := range designs {
		if numCombinations(towels, design, cache) > 0 {
			count++
		}
	}
	fmt.Println(count)
}
