package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"fmt"
)

type Solution struct {
	g grid.Grid
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Solution) Solve() any {
	part1 := 0
	part2 := 0
	R := len(s.g)
	C := len(s.g[0])
	visited := set.New[[2]int]()
	regions := []*set.Set[[2]int]{}
	regionBounds := [][4]int{}
	minr, minc, maxr, maxc := 0, 0, 0, 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			p := [2]int{r, c}
			if visited.Contains(p) {
				continue
			}
			visited.Add(p)
			region := set.New(p)
			minr = min(p[0], R-1)
			minc = min(p[1], C-1)
			maxr = max(p[0], 0)
			maxc = max(p[1], 0)
			q := [][2]int{}
			q = append(q, p)
			for len(q) > 0 {
				pq := q[0]
				q = q[1:]
				for _, pp := range s.g.Neighbours(4, pq[0], pq[1], func(row, col int) bool {
					return s.g[row][col] == s.g[r][c] && !region.Contains([2]int{row, col})
				}) {
					visited.Add(pp)
					region.Add(pp)
					q = append(q, pp)
					minr, minc = min(minr, pp[0]), min(minc, pp[1])
					maxr, maxc = max(maxr, pp[0]), max(maxc, pp[1])
				}
			}
			regions = append(regions, region)
			regionBounds = append(regionBounds, [4]int{minr, minc, maxr, maxc})
		}
	}
	for ri, region := range regions {
		part1 += region.Len() * perimeter(region, s.g)
		part2 += region.Len() * sides(region, regionBounds[ri])
	}
	return [2]int{part1, part2}
}

func perimeter(r *set.Set[[2]int], g grid.Grid) int {
	perimeter := 0
	for p := range r.Values() {
		for _, nb := range g.AllNeighbours(4, p[0], p[1]) {
			if !r.Contains(nb) {
				perimeter++
			}
		}
	}
	return perimeter
}

func pr(g [][]bool, ch string) {
	R, C := len(g), len(g[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] {
				fmt.Print(ch)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func makeRegionGrid(region *set.Set[[2]int], bounds [4]int) [][]bool {
	minr, minc, maxr, maxc := bounds[0], bounds[1], bounds[2], bounds[3]
	R := maxr - minr + 1
	C := maxc - minc + 1
	g := make([][]bool, R)
	for r := 0; r < R; r++ {
		g[r] = make([]bool, C)
	}
	for p := range region.Values() {
		g[p[0]-minr][p[1]-minc] = true
	}
	return g
}

// The region has bounds, which define a grid where we set each position in the
// region to true.
// We do 2 sweeps through the grid:
//  1. tl->br for left- and top edges.
//  2. br->tl for right- and bottom edges.
//
// Each sweep we consider the current (r,c) to be in a 2x2 grid:
// For le and te, we look at (r,c) being at the br in that grid
// For re and be, we look at (r,c) being at the tl in that grid
//
// In either sweep, we have 16 possible states of that 2x2 grid
// { ..  00 01 10 11 00 00 00 01 01 01 10 10 10 11 11 11
// { .#  00 00 00 00 01 10 11 01 10 11 01 10 11 01 10 11
// { le:  0  0  0  0  1  0  0  0  0  0  1  0  0  1  0  0
// { te:  0  0  0  0  1  0  0  0  0  0  1  0  1  0  0  0
//
// Right and bottom edges can be counted by considering the tl as g[r][c]
// { #.  00 01 10 11 00 00 00 01 01 01 10 10 10 11 11 11
// { ..  00 00 00 00 01 10 11 01 10 11 01 10 11 01 10 11
// { re:  0  0  1  0  0  0  0  0  0  0  1  0  1  0  0  0
// { be:  0  0  1  0  0  0  0  0  0  0  1  0  0  1  0  0
//
// So we need to check only a few situations per sweep
// { ..  00 10 10 11
// { .#  01 01 11 01
// { le:  1  1  0  1
// { te:  1  1  1  0
//
// Right and bottom edges can be counted by considering the tl as g[r][c]
// { #.  10 10 10 11
// { ..  00 01 11 01
// { re:  1  1  1  0
// { be:  1  1  0  1
func sides(region *set.Set[[2]int], bounds [4]int) int {
	g := makeRegionGrid(region, bounds)
	R := len(g)
	C := len(g[0])

	sides := 0
	// sweep tl -> br
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			tl, tr, bl, br := false, false, false, g[r][c]
			if r > 0 {
				tr = g[r-1][c]
			}
			if r > 0 && c > 0 {
				tl = g[r-1][c-1]
			}
			if c > 0 {
				bl = g[r][c-1]
			}
			// { ..  00 10 10 11
			// { .#  01 01 11 01
			// { le:  1  1  0  1
			// { te:  1  1  1  0
			if !tl && !tr && !bl && br {
				sides += 2
			}
			if tl && !tr && !bl && br {
				sides += 2
			}
			if tl && !tr && bl && br {
				sides++
			}
			if tl && tr && !bl && br {
				sides++
			}
		}
	}
	// sweep br -> tl
	for r := R - 1; r >= 0; r-- {
		for c := C - 1; c >= 0; c-- {
			tl, tr, bl, br := g[r][c], false, false, false
			if r < R-1 {
				bl = g[r+1][c]
			}
			if r < R-1 && c < C-1 {
				br = g[r+1][c+1]
			}
			if c < C-1 {
				tr = g[r][c+1]
			}
			// { #.  10 10 10 11
			// { ..  00 01 11 01
			// { re:  1  1  1  0
			// { be:  1  1  0  1
			if tl && !tr && !bl && !br {
				sides += 2
			}
			if tl && !tr && !bl && br {
				sides += 2
			}
			if tl && !tr && bl && br {
				sides++
			}
			if tl && tr && !bl && br {
				sides++
			}
		}
	}
	return sides
}

func main() {
	s := &Solution{}
	s.g = grid.FromFile("./input")
	fmt.Println()
	fmt.Println(s.Solve())
}
