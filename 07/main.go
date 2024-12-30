package main

import (
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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

func main() {
	eqs := [][]int{}
	input.Lines(func(s *bufio.Scanner) {
		eqs = append(
			eqs,
			slice.Map(slice.Int, strings.Split(strings.Replace(s.Text(), ":", "", 1), " ")))
	})

	p1 := 0
	p2 := 0
	for _, q := range eqs {
		if Equates(q[0], q[1:], false) {
			p1 += q[0]
		}
		if Equates(q[0], q[1:], true) {
			p2 += q[0]
		}
	}
	fmt.Println(p1, p2)
}
