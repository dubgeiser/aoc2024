package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var re2 = regexp.MustCompile(`^mul\((\d+),(\d+)\)`)

func main() {
	lines := []string{}
	opLines := [][][]string{} // [[[mul(x,y) x y], [mul(a,b) a b],...], ... ]
	input.Lines(func(s *bufio.Scanner) {
		lines = append(lines, s.Text())
		opLines = append(opLines, re.FindAllStringSubmatch(s.Text(), -1))
	})

	p1 := 0
	for _, l := range opLines {
		for _, op := range l {
			n1, _ := strconv.Atoi(op[1])
			n2, _ := strconv.Atoi(op[2])
			p1 += n1 * n2
		}
	}

	p2 := 0
	input := strings.Join(lines, "")
	enabled := true
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i:], "do()") {
			enabled = true
		}
		if strings.HasPrefix(input[i:], "don't()") {
			enabled = false
		}
		if enabled {
			ops := re2.FindAllStringSubmatch(input[i:], -1)
			if len(ops) == 0 {
				continue
			}
			op := ops[0]
			n1, _ := strconv.Atoi(op[1])
			n2, _ := strconv.Atoi(op[2])
			p2 += n1 * n2
		}
	}
	fmt.Println(p1, p2)
}
