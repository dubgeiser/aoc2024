package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

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

func CountUniquePaths(g [][]int, r, c int) int {
	paths := [][][2]int{}
	var path [][2]int
	q := [][3]int{{r, c, g[r][c]}}
	path = append(path, [2]int{r, c})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr[2] == 9 {
			if !Contains(paths, path) {
				paths = append(paths, path)
			}
		}
		for adj := range slices.Values(Adjacents(g, curr[0], curr[1])) {
			q = append(q, adj)
		}
		path = append(path, [2]int{curr[0], curr[1]})
	}
	return len(paths)
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

// Can't use a set.Set, becaue [][][2]int does not implement Comparable
// ... or I did not read the compiler's error message
func Contains(paths [][][2]int, path [][2]int) bool {
	for p := range slices.Values(paths) {
		if slices.Equal(p, path) {
			return true
		}
	}
	return false
}

func main() {
	g := [][]int{}
	trailHeads := [][2]int{}
	input.Lines(func(s *bufio.Scanner) {
		g = append(g, slice.Map(slice.Int,strings.Split(s.Text(), "")))
	})
	R := len(g)
	C := len(g[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == 0 {
				trailHeads = append(trailHeads, [2]int{r, c})
			}
		}
	}
	p1, p2 := 0, 0
	for _, p := range trailHeads {
		p1 += CountUniqueEnds(g, p[0], p[1])
		p2 += CountUniquePaths(g, p[0], p[1])
	}
	fmt.Println(p1, p2)
}
