package main

import (
	"aoc2024/lib/file"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Solution struct {
	l1 []int
	l2 []int
}

func (s *Solution) Part1() int {
	total := 0
	slices.Sort(s.l1)
	slices.Sort(s.l2)
	for i, v := range s.l1 {
		total += int(math.Abs(float64(v - s.l2[i])))
	}
	return total
}

func count(l []int, n int) int {
	total := 0
	for _, v := range l {
		if v == n {
			total++
		}
	}
	return total
}

func (s *Solution) Part2() int {
	total := 0
	for _, n := range s.l1 {
		total += n * count(s.l2, n)
	}
	return total
}

func (s *Solution) ProcessLine(lineIndex int, line string) {
	sNumbers := strings.Split(line, "   ")
	n1, _ := strconv.Atoi(sNumbers[0])
	n2, _ := strconv.Atoi(sNumbers[1])
	s.l1 = append(s.l1, n1)
	s.l2 = append(s.l2, n2)
}

func main() {
	s := &Solution{}
	file.ReadLines("./input", s)
	fmt.Println()
	fmt.Println("Part1:", s.Part1())
	fmt.Println("Part2:", s.Part2())
}
