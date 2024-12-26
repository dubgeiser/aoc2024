package main

import (
	"aoc/lib/file"
	"aoc/lib/grid"
	"fmt"
	"strings"
)

func findStart(g grid.Grid) (int, int) {
	sr, sc := 0, 0
	R := len(g)
	C := len(g[0])
findstart:
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == '@' {
				sr, sc = r, c
				break findstart
			}
		}
	}
	return sr, sc
}

func main() {
	fmt.Println()
	g_, m_ := file.ReadTwoParts(file.Input())
	m := []byte(strings.Replace(m_, "\n", "", -1))
	g := grid.FromString(g_)
	R := len(g)
	C := len(g[0])
	M := len(m)
	m2d := map[byte][2]int{
		'<': {0, -1},
		'>': {0, 1},
		'^': {-1, 0},
		'v': {1, 0},
	}

	r, c := findStart(g)
	for i := 0; i < M; i++ {
		// Collect positions of consecutive boxes, so we can move them all
		// together, if possible.
		consecutiveBoxes := [][2]int{}
		dr, dc := m2d[m[i]][0], m2d[m[i]][1]
		rr, cc := r, c
		move := true
		for {
			rr, cc = rr+dr, cc+dc
			ch := g[rr][cc]
			// Cannot move, bail
			if ch == '#' {
				move = false
				break
			}
			if ch == 'O' {
				consecutiveBoxes = append(consecutiveBoxes, [2]int{rr, cc})
			}
			// Just move, we don't have any more consecutive boxes.
			if ch == '.' {
				break
			}
		}
		if !move {
			continue
		}
		// Move robot and boxes
		g[r][c] = '.'
		g[r+dr][c+dc] = '@'
		r, c = r+dr, c+dc
		for i := 0; i < len(consecutiveBoxes); i++ {
			br, bc := consecutiveBoxes[i][0], consecutiveBoxes[i][1]
			g[br+dr][bc+dc] = 'O'
		}
	}
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] != 'O' {
				continue
			}
			count += 100*r + c
		}
	}

	fmt.Println(count)
}
