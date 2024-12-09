package main

import (
	"aoc/lib/file"
	"fmt"
	"slices"
	"strconv"
)

const FREE int = -1

type Solution struct {
	disk []int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.disk = make([]int, 0)
	fileId := 0
	for i := 0; i < len(line); i++ {
		n, _ := strconv.Atoi(string(line[i]))
		blocks := make([]int, n)
		if i%2 == 0 {
			for j := 0; j < len(blocks); j++ {
				blocks[j] = fileId
			}
			s.disk = slices.Concat(s.disk, blocks)
			fileId++
		} else {
			for j := 0; j < len(blocks); j++ {
				blocks[j] = FREE
			}
			s.disk = slices.Concat(s.disk, blocks)
		}
	}
}

func (s *Solution) Solve() any {
	dfrag := slices.Clone(s.disk)
	// Indices on the disk of the blocks that are free space.
	free := make([]int, 0)
	for i := 0; i < len(dfrag); i++ {
		if dfrag[i] == FREE {
			free = append(free, i)
		}
	}
	for i := 0; i < len(free); i++ {
		for dfrag[len(dfrag)-1] == FREE {
			dfrag = dfrag[:len(dfrag)-1]
		}
		if len(dfrag) > free[i] && dfrag[len(dfrag)-1] != FREE {
			dfrag[free[i]] = dfrag[len(dfrag)-1]
		}
		if len(dfrag) > free[i] {
			dfrag = dfrag[:len(dfrag)-1]
		}
	}
	checksum := 0
	for i, d := range dfrag {
		checksum += i * d
	}
	return [2]int{checksum, 0}
}

func main() {
	s := &Solution{}
	_, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(s.Solve())
}
