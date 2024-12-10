package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"slices"
	"strings"
)

type Solution struct {
	g          [][]int
	trailHeads [][2]int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.g = append(s.g, slice.Map(slice.Int, strings.Split(line, "")))
}

func (s *Solution) Solve() any {
	R := len(s.g)
	C := len(s.g[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if s.g[r][c] == 0 {
				s.trailHeads = append(s.trailHeads, [2]int{r, c})
			}
		}
	}
	p1 := 0
	p2 := 0
	for p := range slices.Values(s.trailHeads) {
		p1 += CountUniqueEnds(s.g, p[0], p[1])
	}
	return [2]int{p1, p2}
}

// Based on BFS, Algo Book p. 556
// Queue (fifo) is implemented with a slice
// Enqueue: q = append(q, p)
// Dequeue: p = q[0]; q = q[1:]
func CountUniqueEnds(g [][]int, r, c int) int {
	visited := set.New[[2]int]()
	ends := set.New[[2]int]()
	q := [][3]int{}
	q = append(q, [3]int{r, c, g[r][c]})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr[2] == 9 {
			ends.Add([2]int{curr[0], curr[1]})
		}
		if visited.Contains([2]int{curr[0], curr[1]}) {
			continue
		}
		for adj := range slices.Values(Adjacents(g, curr[0], curr[1])) {
			q = append(q, adj)
		}
		visited.Add([2]int{curr[0], curr[1]})
	}
	return ends.Len()
}

var dirs = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func Adjacents(g [][]int, r, c int) [][3]int {
	R := len(g)
	C := len(g[0])
	adj := [][3]int{}
	for _, p := range dirs {
		rr, cc := r+p[0], c+p[1]
		if rr < 0 || cc < 0 || rr >= R || cc >= C {
			continue
		}
		if g[r][c] == g[rr][cc]-1 {
			adj = append(adj, [3]int{rr, cc, g[rr][cc]})
		}
	}
	return adj
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
