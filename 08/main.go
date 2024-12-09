package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"fmt"
)

type Solution struct {
	g   grid.Grid
	a2p map[byte][][2]int // antenna (char) -> [all positions it occurs]
}

func (s *Solution) Solve() any {
	antinodes := set.New[[2]int]()
	antinodes2 := set.New[[2]int]()
	s.MapAntennas()
	R := len(s.g)
	C := len(s.g[0])
	for _, pa := range s.a2p {
		for i := 0; i < len(pa)-1; i++ {
			for j := i + 1; j < len(pa); j++ {
				r1 := pa[i][0]
				c1 := pa[i][1]
				r2 := pa[j][0]
				c2 := pa[j][1]
				rr1 := r1 + r1 - r2
				cc1 := c1 + c1 - c2
				rr2 := r2 + r2 - r1
				cc2 := c2 + c2 - c1
				if rr1 >= 0 && rr1 < R && cc1 >= 0 && cc1 < C {
					antinodes.Add([2]int{rr1, cc1})
				}
				if rr2 >= 0 && rr2 < R && cc2 >= 0 && cc2 < C {
					antinodes.Add([2]int{rr2, cc2})
				}

				dr := r2 - r1
				dc := c2 - c1
				r := r1
				c := c1
				for r >= 0 && r < R && c >= 0 && c < C {
					antinodes2.Add([2]int{r, c})
					r += dr
					c += dc
				}
				r = r2
				c = c2
				for r >= 0 && r < R && c >= 0 && c < C {
					antinodes2.Add([2]int{r, c})
					r -= dr
					c -= dc
				}
			}
		}
	}
	return [2]int{antinodes.Len(), antinodes2.Len()}
}

func (s *Solution) MapAntennas() {
	s.a2p = make(map[byte][][2]int)
	for r, row := range s.g {
		for c, ch := range row {
			if ch != '.' {
				if _, exists := s.a2p[ch]; !exists {
					s.a2p[ch] = make([][2]int, 0)
				}
				p := [2]int{r, c}
				s.a2p[ch] = append(s.a2p[ch], p)
			}
		}
	}
}

func main() {
	s := &Solution{}
	s.g = grid.FromFile("./input")

	fmt.Println()
	fmt.Println(s.Solve())
}
