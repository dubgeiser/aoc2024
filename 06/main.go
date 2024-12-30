package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/input"
	"fmt"
)

func main() {
	g, R, C := input.Grid()
	sr, sc := 0, 0
all:
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == '^' {
				sr, sc = r, c
				break all
			}
		}
	}
	visited := *set.New[[2]int]()
	r, c := sr, sc
	// 0:up, 1:right, 2:down, 3:left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	di := 0 // start facing up
	for {
		visited.Add([2]int{r, c})
		rr := r + dirs[di][0]
		cc := c + dirs[di][1]
		inBounds := rr >= 0 && rr < R && cc >= 0 && cc < C
		if !inBounds {
			break
		} else if g[rr][cc] == '#' {
			di = (di + 1) % 4
		} else {
			r, c = rr, cc
		}
	}
	fmt.Println(visited.Len())

	result := 0
	visited2 := set.New[[3]int]()
	// We only need to check obstacles placed on the positions the guard can
	// possibly occupy.
	for o := range visited.Values() {
		visited2.Clear()
		r, c := sr, sc
		di := 0 // start facing up
		for {
			// If we've already visited this position in this direction,
			// we have made a loop.
			if visited2.Contains([3]int{r, c, di}) {
				result++
				break
			}
			visited2.Add([3]int{r, c, di})
			rr := r + dirs[di][0]
			cc := c + dirs[di][1]
			inBounds := rr >= 0 && rr < R && cc >= 0 && cc < C
			if !inBounds {
				break
			}
			// If we hit an existing obstacle (or the one we're testing for
			// right now, change direction)
			if g[rr][cc] == '#' || rr == o[0] && cc == o[1] {
				di = (di + 1) % 4
			} else {
				r, c = rr, cc
			}
		}
	}
	fmt.Println(result)
}
