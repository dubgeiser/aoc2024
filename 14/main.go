package main

import (
	"aoc/lib/algo"
	"aoc/lib/file"
	"aoc/lib/slice"
	"fmt"
	"regexp"
)

type Solution struct {
	robots [][4]int
}

var re = regexp.MustCompile(`-?\d+`)

func (s *Solution) ProcessLine(i int, line string) {
	s.robots = append(s.robots, [4]int(slice.Map(slice.Int, re.FindAllString(line, -1))))
}

func endPosition(x, y, dx, dy, X, Y, t int) [2]int {
	return [2]int{algo.Mod((x + t*dx), X), algo.Mod((y + t*dy), Y)}
}

func (s *Solution) Solve(X, Y, t int) any {
	p1 := 0
	p2 := 0
	mx := (X - 1) / 2
	my := (Y - 1) / 2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for i := range s.robots {
		p := endPosition(s.robots[i][0], s.robots[i][1], s.robots[i][2], s.robots[i][3], X, Y, t)
		x, y := p[0], p[1]
		if x == mx || y == my {
			continue
		}
		if x < mx && y < my {
			q1++
		} else if x < mx && y > my {
			q2++
		} else if x > mx && y < my {
			q3++
		} else {
			q4++
		}
	}
	p1 = q1 * q2 * q3 * q4
	return [2]int{p1, p2}
}

func main() {
	s := &Solution{}
	fn := "input"
	_, err := file.ReadLines(fn, s)
	if err != nil {
		panic(err)
	}
	t := 100
	X, Y := 101, 103
	if fn == "sample" {
		X, Y = 11, 7
	}
	fmt.Println()
	fmt.Println(s.Solve(X, Y, t))
}
