package main

import (
	"aoc/lib/input"
	"aoc/lib/slice"
	"fmt"
	"math"
	"regexp"
)

var re = regexp.MustCompile(`\d+`)

// a: button (ax, ay)
// b: button (bx, by)
// p: prize  (px, py)
//
// We must find number of presses of each button (na, nb) so that:
//
//	na*a + nb*b = p
//
// Which means that (order of button presses does not matter):
// na*ax + nb*bx = px AND na*ay + nb*by = py
//
// There aren't any different ways to reach p: Either you get there by pressing
// button `a` a couple of times (`na`) and then `b` a couple of times `nb` a
// couple of times or you don't get there... "lineair onafhankelijk"
// (There are no zeroes or negative numbers in the input)
//
// "Stelsel van 2 vergelijkingen met 2 onbekenden"
// "Combinatiemethode"
// https://www.youtube.com/watch?v=eGMcVSVIt1c
//
// Take nb out of equation, so we have na.
// { na*ax + nb*bx = px  (*by)
// { na*ay + nb*by = py  (*bx)
//
// ax*by*na + bx*by*nb = by*px
// ay*bx*na + bx*by*nb= bx*py  (-)
// ---------------------------
// (ax*by - ay*bx)na = by*px - bx*py
// <=> na := (by*px - bx*py)/(ax*by - ay*bx)
//
// Take na out of equation, so we have nb.
// { na*ax + nb*bx = px (*ay)
// { na*ay + nb*by = py (*ax)
//
// ax*ay*na + ay*bx*nb = ay*px
// ax*ay*na + ax*by*nb= ax*py   (-)
// ----------------------------
// ay*bx*nb - ax*by*nb = ay*px - ax*py
// <=> nb := (ay*px - ax*py)/(ay*bx - ax*by)
func main() {
	games := [][3][2]int{}
	input.Blocks(func(block string) {
		d := slice.Map(slice.Int, re.FindAllString(block, -1))
		fmt.Println(block)
		games = append(games, [3][2]int{
			{d[0], d[1]},  // button a
			{d[2], d[3]},  // button b
			{d[4], d[5]}}) // prize
	})

	p1 := 0
	p2 := 0
	for _, g := range games {
		ax, ay, bx, by, px, py := g[0][0], g[0][1], g[1][0], g[1][1], g[2][0], g[2][1]
		na := float64(by*px-bx*py) / float64(ax*by-ay*bx)
		nb := float64(ay*px-ax*py) / float64(ay*bx-ax*by)

		// Only consider whole numbers, since we probably cannot partly push
		if math.Trunc(na) == na && math.Trunc(nb) == nb && na >= 0 && na <= 100 && nb >= 0 && nb <= 100 {
			p1 += int(na*3 + nb)
		}

		// part 2
		px += 10000000000000
		py += 10000000000000
		na = float64(by*px-bx*py) / float64(ax*by-ay*bx)
		nb = float64(ay*px-ax*py) / float64(ay*bx-ax*by)
		if math.Trunc(na) == na && math.Trunc(nb) == nb && na >= 0 && nb >= 0 {
			p2 += int(na*3 + nb)
		}
	}
	fmt.Println(35082, 82570698600470)
	fmt.Println(p1, p2)
}
