package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func count(l []int, n int) int {
	total := 0
	for _, v := range l {
		if v == n {
			total++
		}
	}
	return total
}

func main() {
	l1, l2 := []int{}, []int{}
	input.Lines(func(s *bufio.Scanner) {
		sNumbers := strings.Split(s.Text(), "   ")
		n1, _ := strconv.Atoi(sNumbers[0])
		n2, _ := strconv.Atoi(sNumbers[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	})
	slices.Sort(l1)
	slices.Sort(l2)
	p1 := 0
	for i, v := range l1 {
		p1 += int(math.Abs(float64(v - l2[i])))
	}

	p2 := 0
	for _, n := range l1 {
		p2 += n * count(l2, n)
	}
	fmt.Println(p1, p2)
}
