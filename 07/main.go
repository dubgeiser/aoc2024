package main

import (
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Solution struct {
	// Using slice of slices instead of map, because I'm not sure if an equation
	// result appears more than once.
	q [][]int
}

func (s *Solution) ProcessLine(i int, line string) {
	line = strings.Replace(line, ":", "", 1)
	qs := strings.Split(line, " ")
	s.q = append(s.q, slice.Map(slice.Int, qs))
}

func (s *Solution) Part1() any {
	result := 0
	for _, q := range s.q {
		if Equates(q[0], q[1:], false) {
			result += q[0]
		}
	}
	return result
}

func Equates(r int, n []int, useConcat bool) bool {
	// End recursion, we have consumed all elements, check the end result
	if len(n) == 1 {
		return r == n[0]
	}
	if Equates(r, slices.Concat([]int{n[0] + n[1]}, n[2:]), useConcat) {
		return true
	}
	if Equates(r, slices.Concat([]int{n[0] * n[1]}, n[2:]), useConcat) {
		return true
	}
	if useConcat {
		c, _ := strconv.Atoi(strconv.Itoa(n[0]) + strconv.Itoa(n[1]))
		if Equates(r, slices.Concat([]int{c}, n[2:]), useConcat) {
			return true
		}
	}
	return false
}

func (s *Solution) Part2() any {
	result := 0
	for _, q := range s.q {
		if Equates(q[0], q[1:], true) {
			result += q[0]
		}
	}
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
