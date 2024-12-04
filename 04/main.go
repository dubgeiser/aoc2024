package main

import (
	"aoc/lib/grid"
	"fmt"
)

type Solution struct {
	G grid.Grid
}

func (s *Solution) Part1() any {
	total := 0
	R := len(s.G)
	C := len(s.G[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// left -> right
			if c+3 < C && s.G[r][c] == 'X' && s.G[r][c+1] == 'M' && s.G[r][c+2] == 'A' && s.G[r][c+3] == 'S' {
				total++
			}
			// right -> left
			if c-3 >= 0 && s.G[r][c] == 'X' && s.G[r][c-1] == 'M' && s.G[r][c-2] == 'A' && s.G[r][c-3] == 'S' {
				total++
			}
			// up -> down
			if r+3 < R && s.G[r][c] == 'X' && s.G[r+1][c] == 'M' && s.G[r+2][c] == 'A' && s.G[r+3][c] == 'S' {
				total++
			}
			// down -> up
			if r+3 < R && s.G[r][c] == 'S' && s.G[r+1][c] == 'A' && s.G[r+2][c] == 'M' && s.G[r+3][c] == 'X' {
				total++
			}
			// left top -> right bot
			if r+3 < R && c+3 < C && s.G[r][c] == 'X' && s.G[r+1][c+1] == 'M' && s.G[r+2][c+2] == 'A' && s.G[r+3][c+3] == 'S' {
				total++
			}
			// right top -> left bot
			if r+3 < R && c-3 >= 0 && s.G[r][c] == 'X' && s.G[r+1][c-1] == 'M' && s.G[r+2][c-2] == 'A' && s.G[r+3][c-3] == 'S' {
				total++
			}
			// left bot -> right top
			if r-3 >= 0 && c+3 < C && s.G[r][c] == 'X' && s.G[r-1][c+1] == 'M' && s.G[r-2][c+2] == 'A' && s.G[r-3][c+3] == 'S' {
				total++
			}
			// right bot -> left top
			if r-3 >= 0 && c-3 >= 0 && s.G[r][c] == 'X' && s.G[r-1][c-1] == 'M' && s.G[r-2][c-2] == 'A' && s.G[r-3][c-3] == 'S' {
				total++
			}
		}
	}
	return total
}

func (s *Solution) Part2() any {
	total := 0
	R := len(s.G)
	C := len(s.G[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// sentinel: if middle is not A, disregard
			if r+1 < R && c+1 < C && s.G[r+1][c+1] != 'A' {
				continue
			}
			// M M
			//  A
			// S S
			if r+2 < R && c+2 < C && s.G[r][c] == 'M' && s.G[r][c+2] == 'M' && s.G[r+2][c] == 'S' && s.G[r+2][c+2] == 'S' {
				total++
			}
			// M S
			//  A
			// M S
			if r+2 < R && c+2 < C && s.G[r][c] == 'M' && s.G[r][c+2] == 'S' && s.G[r+2][c] == 'M' && s.G[r+2][c+2] == 'S' {
				total++
			}
			// S M
			//  A
			// S M
			if r+2 < R && c+2 < C && s.G[r][c] == 'S' && s.G[r][c+2] == 'M' && s.G[r+2][c] == 'S' && s.G[r+2][c+2] == 'M' {
				total++
			}
			// S S
			//  A
			// M M
			if r+2 < R && c+2 < C && s.G[r][c] == 'S' && s.G[r][c+2] == 'S' && s.G[r+2][c] == 'M' && s.G[r+2][c+2] == 'M' {
				total++
			}
		}
	}
	return total
}

func main() {
	s := &Solution{}
	s.G = grid.FromFile("./input")
	fmt.Println()
	fmt.Println("Part1:", s.Part1())
	fmt.Println("Part2:", s.Part2())
}
