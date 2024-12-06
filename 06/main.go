package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/file"
	"aoc/lib/grid"
	"fmt"
	"slices"
)

type Solution struct {
	g  grid.Grid
	sr int
	sc int
}

func (s *Solution) ProcessLine(i int, line string) {
	row := []byte(line)
	c := slices.Index(row, '^')
	if c > -1 {
		row[c] = '.'
		s.sr = i
		s.sc = c
	}
	s.g = append(s.g, row)
}

func (s *Solution) Part1() any {
	visited := set.New[[2]int]()
	r, c := s.sr, s.sc
	R := len(s.g)
	C := len(s.g[0])
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
		} else if s.g[rr][cc] == '#' {
			di = (di + 1) % 4
		} else {
			r, c = rr, cc
		}
	}
	return visited.Len()
}

func (s *Solution) Part2() any {
	result := 0
	visited := set.New[[3]int]()
	R := len(s.g)
	C := len(s.g[0])
	// 0:up, 1:right, 2:down, 3:left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			r, c := s.sr, s.sc
			di := 0 // start facing up
			visited.Clear()
			for {
				// If we've already visited this position in this direction,
				// we have made a loop.
				if visited.Has([3]int{r, c, di}) {
					result++
					break
				}
				visited.Add([3]int{r, c, di})
				rr := r + dirs[di][0]
				cc := c + dirs[di][1]
				inBounds := rr >= 0 && rr < R && cc >= 0 && cc < C
				if !inBounds {
					break
				}
				// If we hit an existing obstacle (or the one we're testing for
				// right now, change direction)
				if s.g[rr][cc] == '#' || rr == i && cc == j {
					di = (di + 1) % 4
				} else {
					r, c = rr, cc
				}
			}
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
