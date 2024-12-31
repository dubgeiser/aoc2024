package main

import (
	"aoc/lib/grid"
	"aoc/lib/input"
	"bufio"
	"fmt"
	"slices"
)

func findStart(g grid.Grid) (int, int) {
	sr, sc := 0, 0
	R := len(g)
	C := len(g[0])
findstart:
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == '@' {
				sr, sc = r, c
				break findstart
			}
		}
	}
	return sr, sc
}

func main() {
	fmt.Println()
	m := []byte{}
	// Keep the original grid, so we can start from that in part 2
	og := grid.Grid{}
	g := grid.Grid{}
	input.TwoParts(
		func(s *bufio.Scanner) {
			g = append(g, slices.Clone(s.Bytes()))
			og = append(og, slices.Clone(s.Bytes()))
		},
		func(s *bufio.Scanner) { m = slices.Concat(m, s.Bytes()) })
	R := len(g)
	C := len(g[0])
	M := len(m)
	m2d := map[byte][2]int{
		'<': {0, -1},
		'>': {0, 1},
		'^': {-1, 0},
		'v': {1, 0},
	}

	r, c := findStart(g)
	for i := 0; i < M; i++ {
		// Collect positions of consecutive boxes, so we can move them all
		// together, if possible.
		consecutiveBoxes := [][2]int{}
		dr, dc := m2d[m[i]][0], m2d[m[i]][1]
		rr, cc := r, c
		move := true
		for {
			rr, cc = rr+dr, cc+dc
			ch := g[rr][cc]
			// Cannot move, bail
			if ch == '#' {
				move = false
				break
			}
			if ch == 'O' {
				consecutiveBoxes = append(consecutiveBoxes, [2]int{rr, cc})
			}
			// Just move, we don't have any more consecutive boxes.
			if ch == '.' {
				break
			}
		}
		if !move {
			continue
		}
		// Move robot and boxes
		g[r][c] = '.'
		g[r+dr][c+dc] = '@'
		r, c = r+dr, c+dc
		for i := 0; i < len(consecutiveBoxes); i++ {
			br, bc := consecutiveBoxes[i][0], consecutiveBoxes[i][1]
			g[br+dr][bc+dc] = 'O'
		}
	}
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] != 'O' {
				continue
			}
			count += 100*r + c
		}
	}
	fmt.Println(count)

	// --[ Part 2 ]--
	// consecutiveBoxes doesn't cut it, we need to take into account:
	// .......
	// .[][]..
	// ..[]...
	// ...@...
	//When going ^ this will push all three boxes upwards, the identical
	// situation exists when pushing downwards.
	// We can track this by fanning out from the position of the robot in the
	// direction that it is moving and collect all '[]' pairs.
	// This whole block can't be moved if only _one_ of the outer boxes can't.
	g, R, C = widenGrid(og)
	r, c = findStart(g)
	for i := 0; i < M; i++ {
		// Start with the position of the robot for tracking, this will ensure
		// that we collect all the connected boxes that can be moved from that
		// position.
		// It also makes our loop for collecting the connected boxes easier.
		btrack := [][2]int{{r, c}}
		dr, dc := m2d[m[i]][0], m2d[m[i]][1]
		rr, cc := r, c
		move := true
		for j := 0; j < len(btrack); j++ {
			// Fanning out
			// Don't track twice, since we're adding the pairs automatically,
			// this would result in a _lot_ of duplicates!
			// Ask me how I know ;-)
			rr, cc = btrack[j][0]+dr, btrack[j][1]+dc
			if slices.Contains(btrack, [2]int{rr, cc}) {
				continue
			}
			ch := g[rr][cc]
			if ch == '#' {
				move = false
				break
			}
			if ch == '[' {
				btrack = append(btrack, [2]int{rr, cc})
				btrack = append(btrack, [2]int{rr, cc + 1})
			}
			if ch == ']' {
				btrack = append(btrack, [2]int{rr, cc})
				btrack = append(btrack, [2]int{rr, cc - 1})
			}
		}
		if !move {
			continue
		}

		// Robot moved: Update the grid
		gg := g.Copy()
		B := len(btrack)
		// Clear all tiles that we have tracked
		for b := 0; b < B; b++ {
			br, bc := btrack[b][0], btrack[b][1]
			g[br][bc] = '.'
		}
		// Then "move" the boxes
		for b := 0; b < B; b++ {
			br, bc := btrack[b][0], btrack[b][1]
			g[br+dr][bc+dc] = gg[br][bc]
		}
		// Last, move the robot
		// I found that the robot's position was sometimes overwritten when
		// moving the boxes.  This did not really impact the logic of the code,
		// but when printing out the grid, it was confusing not to see the robot
		// where it was supposed to be.
		g[r][c] = '.'
		g[r+dr][c+dc] = '@'
		r, c = r+dr, c+dc
	}
	count = 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] != '[' {
				continue
			}
			count += 100*r + c
		}
	}
	fmt.Println(count)
}

func widenGrid(g grid.Grid) (gg grid.Grid, R, C int) {
	R, C = len(g), len(g[0])
	gg = grid.Grid{}
	grow := map[byte]string{'#': "##", '.': "..", 'O': "[]", '@': "@."}
	for r := 0; r < R; r++ {
		gg = append(gg, []byte{})
		for c := 0; c < C; c++ {
			for _, ch := range grow[g[r][c]] {
				gg[r] = append(gg[r], byte(ch))
			}
		}
	}
	R, C = len(gg), len(gg[0])
	return
}
