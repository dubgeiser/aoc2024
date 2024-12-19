package main

import (
	"aoc/lib/file"
	"fmt"
)

type Solution struct {
}

func (s *Solution) ProcessLine(i int, line string) {
}

func main() {
	fmt.Println()
	content, err := os.ReadFile("./input")
	if err != nil {
		panic("Cannot read input!")
	}
	lines := strings.Split(string(content), "\n")
}
