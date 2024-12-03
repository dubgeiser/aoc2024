package main

import (
	"aoc/lib/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var re2 = regexp.MustCompile(`^mul\((\d+),(\d+)\)`)

type Solution struct {
	lines   []string
	opLines [][][]string // [[[mul(x,y) x y], [mul(a,b) a b],...], ... ]
}

func (s *Solution) ProcessLine(i int, line string) {
	s.lines = append(s.lines, line)
	s.opLines = append(s.opLines, re.FindAllStringSubmatch(line, -1))
}

func (s *Solution) Part1() any {
	total := 0
	for _, l := range s.opLines {
		for _, op := range l {
			n1, _ := strconv.Atoi(op[1])
			n2, _ := strconv.Atoi(op[2])
			total += n1 * n2
		}
	}
	return total
}

func (s *Solution) Part2() any {
	total := 0
	input := strings.Join(s.lines, "")
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
			total += n1 * n2
		}
	}
	return total
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
