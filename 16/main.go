package main

import (
	"aoc/lib/algo"
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"aoc/lib/input"
	"container/heap"
	"fmt"
	"slices"
)

type Tile struct {
	r, c  int
	dir   [2]int
	cost  int
	index int
	path  [][2]int
}

type VisitedTile struct {
	r, c int
	dir  [2]int
}

func (i *Tile) convert() VisitedTile {
	return VisitedTile{r: i.r, c: i.c, dir: i.dir}
}

type PriorityQueue []*Tile

func NewPq(i *Tile) *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, i)
	return pq
}

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
	item := x.(*Tile)
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

// Directions, ordered clockwise: NESW
const N, E, S, W = 0, 1, 2, 3

var DCW = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Neighbours(g grid.Grid, i *Tile) []*Tile {
	nbs := []*Tile{}
	dcw := DCW[(algo.Mod(slices.Index(DCW, i.dir)+1, len(DCW)))]
	dccw := DCW[(algo.Mod(slices.Index(DCW, i.dir)-1, len(DCW)))]
	path := slices.Clone(i.path)

	// Forward, apparently, you can't go out of bounds.
	rr, cc := i.r+i.dir[0], i.c+i.dir[1]
	if g[rr][cc] != '#' {
		nbs = append(nbs, &Tile{
			cost: i.cost + 1,
			r:    rr,
			c:    cc,
			dir:  i.dir,
			path: append(path, [2]int{i.r, i.c}),
		})
	}

	// Turn clockwise and counter clockwise
	nbs = append(nbs, &Tile{cost: i.cost + 1000, r: i.r, c: i.c, dir: dcw, path: path})
	nbs = append(nbs, &Tile{cost: i.cost + 1000, r: i.r, c: i.c, dir: dccw, path: path})

	return nbs
}

func GetStart(g grid.Grid, dir [2]int) *Tile {
	R := len(g)
	C := len(g[0])
	sr, sc := -1, -1
all:
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if g[r][c] == 'S' {
				sr = r
				sc = c
				break all
			}
		}
	}
	i := &Tile{cost: 0, r: sr, c: sc, dir: dir, path: [][2]int{{}}}
	return i
}

func main() {
	fmt.Println()
	g, _, _ := input.Grid()

	// Save for part 2.
	cheapest := -1
	start := GetStart(g, DCW[E])

	// --[ Part 1 ]--
	visited := set.New[VisitedTile]()
	pq := NewPq(start)
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Tile)
		if g[curr.r][curr.c] == 'E' {
			cheapest = curr.cost
			fmt.Println(curr.cost)
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

	// --[ Part 2 ]--
	pq = NewPq(start)
	visited = set.New[VisitedTile]()
	places := set.New([2]int{start.r, start.c})
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Tile)
		if curr.cost > cheapest {
			continue
		}
		visited.Add(curr.convert())
		if g[curr.r][curr.c] == 'E' && curr.cost == cheapest {
			for _, place := range curr.path {
				places.Add(place)
			}
		}
		for _, nb := range Neighbours(g, curr) {
			if visited.Contains(nb.convert()) {
				continue
			}
			heap.Push(pq, nb)
		}
	}
	fmt.Println(places.Len())
}
