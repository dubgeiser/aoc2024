package main

import (
	"aoc/lib/input"
	"fmt"
)

func main() {
	g, R, C := input.Grid()
	p1 := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// left -> right
			if c+3 < C && g[r][c] == 'X' && g[r][c+1] == 'M' && g[r][c+2] == 'A' && g[r][c+3] == 'S' {
				p1++
			}
			// right -> left
			if c-3 >= 0 && g[r][c] == 'X' && g[r][c-1] == 'M' && g[r][c-2] == 'A' && g[r][c-3] == 'S' {
				p1++
			}
			// up -> down
			if r+3 < R && g[r][c] == 'X' && g[r+1][c] == 'M' && g[r+2][c] == 'A' && g[r+3][c] == 'S' {
				p1++
			}
			// down -> up
			if r+3 < R && g[r][c] == 'S' && g[r+1][c] == 'A' && g[r+2][c] == 'M' && g[r+3][c] == 'X' {
				p1++
			}
			// left top -> right bot
			if r+3 < R && c+3 < C && g[r][c] == 'X' && g[r+1][c+1] == 'M' && g[r+2][c+2] == 'A' && g[r+3][c+3] == 'S' {
				p1++
			}
			// right top -> left bot
			if r+3 < R && c-3 >= 0 && g[r][c] == 'X' && g[r+1][c-1] == 'M' && g[r+2][c-2] == 'A' && g[r+3][c-3] == 'S' {
				p1++
			}
			// left bot -> right top
			if r-3 >= 0 && c+3 < C && g[r][c] == 'X' && g[r-1][c+1] == 'M' && g[r-2][c+2] == 'A' && g[r-3][c+3] == 'S' {
				p1++
			}
			// right bot -> left top
			if r-3 >= 0 && c-3 >= 0 && g[r][c] == 'X' && g[r-1][c-1] == 'M' && g[r-2][c-2] == 'A' && g[r-3][c-3] == 'S' {
				p1++
			}
		}
	}
	fmt.Println(p1)

	p2 := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			// sentinel: if middle is not A, disregard
			if r+1 < R && c+1 < C && g[r+1][c+1] != 'A' {
				continue
			}
			// M M
			//  A
			// S S
			if r+2 < R && c+2 < C && g[r][c] == 'M' && g[r][c+2] == 'M' && g[r+2][c] == 'S' && g[r+2][c+2] == 'S' {
				p2++
			}
			// M S
			//  A
			// M S
			if r+2 < R && c+2 < C && g[r][c] == 'M' && g[r][c+2] == 'S' && g[r+2][c] == 'M' && g[r+2][c+2] == 'S' {
				p2++
			}
			// S M
			//  A
			// S M
			if r+2 < R && c+2 < C && g[r][c] == 'S' && g[r][c+2] == 'M' && g[r+2][c] == 'S' && g[r+2][c+2] == 'M' {
				p2++
			}
			// S S
			//  A
			// M M
			if r+2 < R && c+2 < C && g[r][c] == 'S' && g[r][c+2] == 'S' && g[r+2][c] == 'M' && g[r+2][c+2] == 'M' {
				p2++
			}
		}
	}
	fmt.Println(p2)
}
