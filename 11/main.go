package main

import (
	"aoc/lib/file"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Solution struct {
	stones []string
}

func (s *Solution) ProcessLine(i int, line string) {
	s.stones = strings.Split(line, " ")
}

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

func (s *Solution) Solve() any {
	stones := slices.Clone(s.stones)
	for i := 0; i < 25; i++ {
		stones = Blink(stones)
	}
	return len(stones)
}

func main() {
	s := &Solution{}
	_, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(s.Solve())
}
