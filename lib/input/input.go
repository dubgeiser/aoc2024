// Package input is the goto place for handling AoC puzzle input.
package input

import (
	"aoc/lib/grid"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Scan the input line per line, each time a new line is scanned, pass the
// scanner to the line processor, so it can scan the current line as it see fit.
func Lines(processor func(s *bufio.Scanner)) {
	in := input()
	defer in.Close()
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		processor(scanner)
	}
}

// Return the input as a Grid and its size.
// Don't do any checking, assume that the input is a valid grid, meaning that,
// since `grid.Grid` is actually `[][]byte`, that `R` will be the length
// of the grid and `C` will be the length of the first row in the grid.
func Grid() (g grid.Grid, R, C int) {
	g = grid.Grid{}
	Lines(func(scanner *bufio.Scanner) {
		g = append(g, slices.Clone(scanner.Bytes()))
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
	Lines(func(scanner *bufio.Scanner) {
		if len(scanner.Bytes()) == 0 {
			processor = p2
			return
		}
		processor(scanner)
	})
}

// Blocks() reads input in blocks of multiple lines.
// Each block is passed verbatim to the processor as a string.
// Every line in a block will be separated by `"\n"`
func Blocks(processor func(s string)) {
	block := strings.Builder{}
	Lines(func(sc *bufio.Scanner) {
		if len(sc.Bytes()) == 0 {
			processor(strings.TrimSpace(block.String()))
			block.Reset()
		} else {
			block.WriteString(sc.Text())
			block.WriteString("\n")
		}
	})
	// Process last block
	processor(strings.TrimSpace(block.String()))
}

// Return the input as one blob.
// Trim any space pre- or succeeding the content.
func Blob() string {
	in := input()
	defer in.Close()
	raw, err := os.ReadFile(in.Name())
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(raw))
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
