package main

import (
	"aoc2024/lib/file"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Solution struct {
	l1 []int
	l2 []int
}

func (s *Solution) Part1() {
	total := 0
	sort.Slice(s.l1, func(i, j int) bool {
		return s.l1[i] < s.l1[j]
	})
	sort.Slice(s.l2, func(i, j int) bool {
		return s.l2[i] < s.l2[j]
	})
	for i, v := range s.l1 {
		total += int(math.Abs(float64(v - s.l2[i])))
	}
	fmt.Println()
	fmt.Println(total)
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
	s.Part1()
}
