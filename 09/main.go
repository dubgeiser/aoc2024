package main

import (
	"aoc/lib/file"
	"fmt"
	"slices"
	"strconv"
)

const FREE int = -1

type Solution struct {
	disk      []int
	maxFileId int

	// {{pos, size}, ...}
	free [][2]int

	// {fileId: {pos, size}, ...}
	id2pos map[int][2]int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.disk = make([]int, 0)
	s.id2pos = make(map[int][2]int)
	fileId := 0
	for i := 0; i < len(line); i++ {
		n, _ := strconv.Atoi(string(line[i]))
		blocks := make([]int, n)
		if i%2 == 0 {
			for j := 0; j < len(blocks); j++ {
				blocks[j] = fileId
			}
			s.id2pos[fileId] = [2]int{len(s.disk), len(blocks)}
			s.disk = slices.Concat(s.disk, blocks)
			fileId++
		} else {
			// {posInDisk, noBlocks}, ie. {len(disk), len(blocks)}
			s.free = append(s.free, [2]int{len(s.disk), len(blocks)})
			for j := 0; j < len(blocks); j++ {
				blocks[j] = FREE
			}
			s.disk = slices.Concat(s.disk, blocks)
		}
		s.maxFileId = fileId - 1
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
	p1 := CheckSum(dfrag)

	// part 2
	dfrag = slices.Clone(s.disk)
	id := s.maxFileId
	for id > 0 {
		for dfrag[len(dfrag)-1] == FREE {
			dfrag = dfrag[:len(dfrag)-1]
		}
		iFile := s.id2pos[id][0]
		nFile := s.id2pos[id][1]
		for i := 0; i < len(s.free); i++ {
			// pr(dfrag)
			iFree := s.free[i][0]
			nFree := s.free[i][1]
			if iFree >= iFile {
				s.free = s.free[:i]
				break
			}
			if nFree >= nFile {
				for n := 0; n < nFile; n++ {
					dfrag[n+iFree] = dfrag[n+iFile]
					dfrag[n+iFile] = FREE
				}
				if nFree == nFile {
					s.free = slices.Delete(s.free, i, i+1)
				} else {
					s.free[i] = [2]int{iFree + nFile, nFree - nFile}
				}
				break
			}
		}
		id--
	}
	p2 := CheckSum(dfrag)
	return [2]int{p1, p2}
}

func CheckSum(disk []int) int {
	checksum := 0
	for i, id := range disk {
		if id == FREE {
			continue
		}
		checksum += i * id
	}
	return checksum
}

func pr(disk []int) {
	for id := range slices.Values(disk) {
		c := strconv.Itoa(id)
		if id == FREE {
			c = "."
		}
		fmt.Print(c)
	}
	fmt.Println()
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
