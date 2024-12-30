package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Blink(stones []string) []string {
	result := []string{}
	for stone := range slices.Values(stones) {
		if stone == "0" {
			result = append(result, "1")
		} else {
			l := len(stone)
			if l%2 == 0 {
				s1, s2 := stone[:l/2], stone[l/2:]
				n1, _ := strconv.Atoi(s1)
				n2, _ := strconv.Atoi(s2)
				result = append(result, strconv.Itoa(n1))
				result = append(result, strconv.Itoa(n2))
			} else {
				n, _ := strconv.Atoi(stone)
				n *= 2024
				result = append(result, strconv.Itoa(n))
			}
		}
	}
	return result
}

type StoneCount struct {
	stone string
	steps int
}

// We're gonna have a ton of 0's, 1's, 2024's and more:
// For each combination of {stone, steps}, cache the result.
func BlinkFasterJoeyTribiani(stone string, steps int, cache map[StoneCount]int) int {
	if steps == 0 {
		return 1
	}
	cacheKey := StoneCount{stone, steps}
	if n, ok := cache[cacheKey]; ok {
		return n
	}
	result := 0
	if stone == "0" {
		result = BlinkFasterJoeyTribiani("1", steps-1, cache)
	} else {
		l := len(stone)
		if l%2 == 0 {
			s1, s2 := stone[:l/2], stone[l/2:]
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			result = BlinkFasterJoeyTribiani(strconv.Itoa(n1), steps-1, cache)
			result += BlinkFasterJoeyTribiani(strconv.Itoa(n2), steps-1, cache)
		} else {
			n, _ := strconv.Atoi(stone)
			n *= 2024
			result = BlinkFasterJoeyTribiani(strconv.Itoa(n), steps-1, cache)
		}
	}
	cache[cacheKey] = result
	return result
}

func main() {
	stones := []string{}
	input.Lines(func(s *bufio.Scanner) { stones = strings.Split(s.Text(), " ") })
	ss := slices.Clone(stones)
	for i := 0; i < 25; i++ {
		ss = Blink(ss)
	}
	fmt.Println(len(ss))

	p2 := 0
	for stone := range slices.Values(stones) {
		p2 += BlinkFasterJoeyTribiani(stone, 75, make(map[StoneCount]int))
	}
	fmt.Println(p2)
}
