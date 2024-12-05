package main

import (
	"aoc/lib/file"
	"fmt"
)

type Solution struct {
}

func (s *Solution) ProcessLine(i int, line string) {
}

func (s *Solution) Part1() any {
	result := 0
	return result
}

func (s *Solution) Part2() any {
	result := 0
	return result
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
