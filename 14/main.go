package main

import (
	"aoc/lib/algo"
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"fmt"
	"math"
	"regexp"
)

var re = regexp.MustCompile(`-?\d+`)

func endPosition(x, y, dx, dy, X, Y, t int) [2]int {
	return [2]int{algo.Mod((x + t*dx), X), algo.Mod((y + t*dy), Y)}
}

func MoveRobots(robots [][4]int, X, Y, t int) []int {
	mx := (X - 1) / 2
	my := (Y - 1) / 2
	q := []int{0, 0, 0, 0}
	for i := range robots {
		p := endPosition(robots[i][0], robots[i][1], robots[i][2], robots[i][3], X, Y, t)
		x, y := p[0], p[1]
		if x == mx || y == my {
			continue
		}
		if x < mx && y < my {
			q[0]++
		} else if x < mx && y > my {
			q[1]++
		} else if x > mx && y < my {
			q[2]++
		} else {
			q[3]++
		}
	}
	return q
}

func Product(nums ...int) int {
	p := 1
	for _, n := range nums {
		p *= n
	}
	return p
}

func pr(robots [][4]int, X, Y, t int) {
	g := make([][]byte, X)
	for x := 0; x < X; x++ {
		g[x] = make([]byte, Y)
		for y := 0; y < Y; y++ {
			g[x][y] = ' '
		}
	}
	for i := range robots {
		p := endPosition(robots[i][0], robots[i][1], robots[i][2], robots[i][3], X, Y, t)
		g[p[0]][p[1]] = '*'
	}

	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			fmt.Print(string(g[x][y]))
		}
		fmt.Println()
	}
}

func main() {
	robots := [][4]int{}
	input.Lines(func(s *bufio.Scanner) {
		robots = append(robots, [4]int(slice.Map(slice.Int, re.FindAllString(s.Text(), -1))))
	})
	X, Y, t := 101, 103, 100
	p1 := Product(MoveRobots(robots, X, Y, t)...)

	p2 := 0
	minq := math.MaxInt
	for i := 0; i < 6621; i++ {
		q := Product(MoveRobots(robots, X, Y, i)...)
		if q < minq {
			minq = q
			p2 = i
			// fmt.Println(p2)
		}
	}

	// A Christmas tree!
	// pr(robots, X, Y, p2)
	fmt.Println(p1, p2)
}
