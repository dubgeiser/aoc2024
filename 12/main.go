package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"fmt"
)

type Solution struct {
	g grid.Grid
}

func (s *Solution) Solve() any {
	part1 := 0
	part2 := 0
	R := len(s.g)
	C := len(s.g[0])
	visited := set.New[[2]int]()
	regions := []*set.Set[[2]int]{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			p := [2]int{r, c}
			if visited.Contains(p) {
				continue
			}
			visited.Add(p)
			region := set.New(p)
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
				}
			}
			regions = append(regions, region)
		}
	}
	for _, region := range regions {
		part1 += region.Len() * perimeter(region, s.g)
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

func main() {
	s := &Solution{}
	s.g = grid.FromFile("./input")
	fmt.Println()
	fmt.Println(s.Solve())
}
