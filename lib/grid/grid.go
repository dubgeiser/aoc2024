package grid

import (
	"aoc/lib/file"
	"fmt"
	"strings"
)

type Grid [][]byte

var DIRS = map[int][][2]int{
	4: {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	8: {{0, -1}, {0, 1}, {1, 0}, {1, -1}, {1, 1}, {-1, 0}, {-1, -1}, {-1, 1}},
}

// Return the neighbours in the grid for a given position `r,c`
// Only return in-bounds positions for which `fn(r,c)` is `true`
func (g Grid) Neighbours(numDirs, r, c int, fn func(row, col int) bool) [][2]int {
	dirs := DIRS[numDirs]
	D := len(dirs)
	R := len(g)
	C := len(g[0])
	nbs := [][2]int{}
	for i := 0; i < D; i++ {
		rr, cc := r+dirs[i][0], c+dirs[i][1]
		if rr >= 0 && rr < R && cc >= 0 && cc < C && fn(rr, cc) {
			nbs = append(nbs, [2]int{rr, cc})
		}
	}
	return nbs
}

func (g Grid) AllNeighbours(numDirs, r, c int) [][2]int {
	dirs := DIRS[numDirs]
	D := len(dirs)
	nbs := [][2]int{}
	for i := 0; i < D; i++ {
		rr, cc := r+dirs[i][0], c+dirs[i][1]
		nbs = append(nbs, [2]int{rr, cc})
	}
	return nbs
}

type GridBuilder struct {
	g Grid
}

func (b *GridBuilder) ProcessLine(i int, line string) {
	b.g = append(b.g, []byte(line))
}

func FromFile(fn string) Grid {
	b := &GridBuilder{}
	file.ReadLines(fn, b)
	return b.g
}

func FromString(sg string) Grid {
	g := [][]byte{}
	for _, line := range strings.Split(sg, "\n") {
		g = append(g, []byte(line))
	}
	return g
}

func Print(g Grid) {
	for _, row := range g {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}
