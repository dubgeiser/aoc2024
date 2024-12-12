package grid

import (
	"aoc/lib/file"
	"fmt"
)

type Grid [][]byte

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

func Print(g Grid) {
	for _, row := range g {
		for _, c := range row {
			fmt.Printf("%s", string(c))
		}
		fmt.Println()
	}
}
