package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var OPS = map[string]func(x int, y int) int{
	"AND": func(x, y int) int { return x & y },
	"XOR": func(x, y int) int { return x ^ y },
	"OR":  func(x, y int) int { return x | y },
}

func val(w string, wires map[string]int, exp map[string][3]string) int {
	if v, ok := wires[w]; ok {
		return v
	}
	x, op, y := exp[w][0], exp[w][1], exp[w][2]
	wires[w] = OPS[op](val(x, wires, exp), val(y, wires, exp))
	return wires[w]
}

func main() {
	wires := map[string]int{}
	expr := map[string][3]string{}
	input.TwoParts(func(s *bufio.Scanner) {
		w := strings.Split(s.Text(), ": ")
		d, _ := strconv.Atoi(w[1])
		wires[w[0]] = d
	}, func(s *bufio.Scanner) {
		//g -> x, OP, y, z
		g := strings.Split(strings.Replace(s.Text(), " ->", "", 1), " ")
		expr[g[3]] = [3]string{g[0], g[1], g[2]}
	})

	for w := range expr {
		wires[w] = val(w, wires, expr)
	}

	re := regexp.MustCompile(`^z[\d+]`)
	bin := strings.Builder{}
	sorted := slices.Sorted(maps.Keys(wires))
	slices.Reverse(sorted)
	for _, w := range sorted {
		if re.MatchString(w) {
			bin.WriteString(strconv.Itoa(wires[w]))
		}
	}
	d, _ := strconv.ParseInt(bin.String(), 2, 64)
	fmt.Println(d)
}
