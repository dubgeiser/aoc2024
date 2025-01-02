package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/grid"
	"aoc/lib/input"
	"aoc/lib/slice"
	"bufio"
	"container/heap"
	"fmt"
	"regexp"
)

type Move struct {
	r, c  int
	steps int
	index int
}

// {row, col, steps}
type PriorityQueue []*Move

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].steps < pq[j].steps
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	m := x.(*Move)
	m.index = pq.Len()
	*pq = append(*pq, m)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func minSteps(size int, g grid.Grid) int {
	sr, sc := 0, 0
	visited := set.New([2]int{sr, sc})
	pq := &PriorityQueue{}
	heap.Push(pq, &Move{r: sr, c: sc, steps: 0})
	for pq.Len() > 0 {
		m := heap.Pop(pq).(*Move)
		if m.r == size && m.c == size {
			return m.steps
		}
		visited.Add([2]int{m.r, m.c})
		for _, nb := range g.Neighbours(4, m.r, m.c, func(row, col int) bool {
			if g[row][col] == '.' {
				return true
			}
			return false
		}) {
			if visited.Contains(nb) {
				continue
			}
			visited.Add([2]int{nb[0], nb[1]})
			heap.Push(pq, &Move{r: nb[0], c: nb[1], steps: m.steps + 1})
		}
	}
	return -1
}

func main() {
	size := 70     // 6 for sample, 70 for input
	nBytes := 1024 // 12 for sample, 1024 for input
	b := [][2]int{}
	re := regexp.MustCompile(`\d+`)
	input.Lines(func(s *bufio.Scanner) {
		d := slice.Map(slice.Int, re.FindAllString(s.Text(), -1))
		// y = row, x = col
		b = append(b, [2]int{d[1], d[0]})
	})
	g := grid.New(size+1, size+1, '.')

	// --[ Part 1]--
	g.MarkAll(b[:nBytes], '#')
	fmt.Println(minSteps(size, g))

	// --[Part 2]--
	// We could do some bisect approacht to make this faster
	g = grid.New(size+1, size+1, '.')
	for i:=0;i<len(b);i++ {
		g.Mark(b[i], '#')
		if minSteps(size, g) > 0 {
			continue
		}
		// y = row, x = col
		fmt.Printf("%d,%d", b[i][1], b[i][0])
		break
	}
}
