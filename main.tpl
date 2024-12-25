package main

import (
	"aoc/lib/file"
	"aoc/lib/grid"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()

	for _, line := range strings.Split(file.Read(), "\n") {
	}

	g, R, C := grid.FromFile(file.Input())

	top, bottom := file.ReadTwoParts(file.Input())

	content, err := os.ReadFile(file.Input())
	if err != nil {
		panic("Cannot read input!")
	}

}
