package main

import (
	"aoc/lib/input"
	"fmt"
	"strings"
)

func pinHeights(s string) [5]int {
	h := [5]int{}
	rows := strings.Split(s, "\n")[1:]
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if rows[r][c] == '#' {
				h[c]++
			}
		}
	}
	return h
}

func main() {
	// locks top row filled
	// the keys bottom row filled.
	locks, keys := [][5]int{}, [][5]int{}
	input.Blocks(func(s string) {
		h := pinHeights(s)
		if strings.HasPrefix(s, "#####") {
			locks = append(locks, h)
		} else {
			keys = append(keys, h)
		}
	})
	p1 := 0
	for l:=0;l<len(locks);l++{
		for k:=0;k<len(keys);k++ {
			fits := true
			for i:=0;i<5;i++ {
				if locks[l][i] + keys[k][i] > 5 {
					fits = false
				}
			}
			if fits {
				p1++
			}
		}
	}
	fmt.Println(p1)
}
