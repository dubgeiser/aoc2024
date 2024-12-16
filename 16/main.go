package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"container/heap"
	"fmt"
	"slices"
)

type Item struct {
	r, c  int
	dir   [2]int
	cost  int
	index int
}

type VisitedItem struct {
	r, c int
	dir  [2]int
}

func (i *Item) convert() VisitedItem {
	return VisitedItem{r: i.r, c: i.c, dir: i.dir}
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func mod(a, b int) int {
	return (a%b + b) % b
}

// Directions, ordered clockwise: NESW
var DCW = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Neighbours(g grid.Grid, i *Item) []*Item {
	nbs := []*Item{}

	// Forward, apparently, you can't go out of bounds.
	rr, cc := i.r+i.dir[0], i.c+i.dir[1]
	if g[rr][cc] != '#' {
		nbs = append(nbs, &Item{cost: i.cost + 1, r: rr, c: cc, dir: i.dir})
	}

	// Turn clockwise and counter clockwise
	dcw := DCW[(mod(slices.Index(DCW, i.dir)+1, len(DCW)))]
	dccw := DCW[(mod(slices.Index(DCW, i.dir)-1, len(DCW)))]
	nbs = append(nbs, &Item{cost: i.cost + 1000, r: i.r, c: i.c, dir: dcw})
	nbs = append(nbs, &Item{cost: i.cost + 1000, r: i.r, c: i.c, dir: dccw})

	return nbs
}

func main() {
	part1 := 0
	part2 := 0
	g := grid.FromFile("./input")
	R := len(g)
	C := len(g[0])

	// Determine start and endpoint
	var sr, sc, er, ec int
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == 'S' {
				sr = r
				sc = c
			}
			if g[r][c] == 'E' {
				er = r
				ec = c
			}
		}
	}

	visited := set.New[VisitedItem]()
	start := &Item{cost: 0, r: sr, c: sc, dir: [2]int{0, 1}} // Facing East.
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, start)
	visited.Add(start.convert())
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if er == curr.r && ec == curr.c {
			part1 = curr.cost
			break
		}
		visited.Add(curr.convert())
		for _, nb := range Neighbours(g, curr) {
			if visited.Contains(nb.convert()) {
				continue
			}
			heap.Push(pq, nb)
		}
	}

	fmt.Println()
	fmt.Println(part1, part2)
}
