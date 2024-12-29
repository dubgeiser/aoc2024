// Package input is the goto place for handling AoC puzzle input.
package input

import (
	"aoc/lib/grid"
	"bufio"
	"fmt"
	"os"
)

// Return the input as a slice of strings, for custom parsing.
func Lines() []string {
	lines := []string{}
	ProcessLine(func(scanner *bufio.Scanner) {
		lines = append(lines, scanner.Text())
	})
	return lines
}

// Return the input as a Grid and its size.
// Don't do any checking, assume that the input is a valid grid, meaning that,
// since `grid.Grid` is actually `[][]byte`, that `R` will be the length
// of the grid and `C` will be the length of the first row in the grid.
func Grid() (g grid.Grid, R, C int) {
	g = grid.Grid{}
	ProcessLine(func(scanner *bufio.Scanner) {
		g = append(g, scanner.Bytes())
	})
	R = len(g)
	C = len(g[0])
	return
}

// Assume the input consists of two parts, separated by a blank line.
// Process both parts line by line, using 2 different processors (`p1` and
// `p2`) to process both parts separately.
func TwoParts(p1, p2 func(s *bufio.Scanner)) {
	processor := p1
	ProcessLine(func(scanner *bufio.Scanner) {
		if len(scanner.Bytes()) == 0 {
			processor = p2
			return
		}
		processor(scanner)
	})
}

// Scan the input line per line, each time a new line is scanned, pass the
// scanner to the line processor, so it can scan the current line as it see fit.
// It is assumed that `processor` does no altering operations on the scanner
// itself.
func ProcessLine(processor func(s *bufio.Scanner)) {
	in := input()
	defer in.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		processor(scanner)
	}
}

// Determine the input for a puzzle.
// For determining which input to read, following rules apply:
// If there's input on STDIN, read it,
// else if there's a filename supplied, read that
// else read from default file named `input`
//
// Note that this will return a file pointer and it is up to the user code to
// cleanup / Close() it!
// For this reason, we keep this function local to the package.
func input() *os.File {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	// On Mac this detect STDIN whether using `cat FILE |`
	// or `go run main.go < FILE`
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return os.Stdin
	}
	fn := "input"
	if len(os.Args) > 1 {
		fn = os.Args[1]
	}
	file, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file [%s]:\n%s", fn, err.Error()))
	}
	return file
}
