package main

import (
	"aoc/lib/input"
	"bufio"
	"fmt"
	"maps"
	"strconv"
)

func newSecret(n int) int {
	n = (n ^ (n * 64)) % 16777216
	n = (n ^ (n / 32)) % 16777216
	n = (n ^ (n * 2048)) % 16777216
	return n
}

func Lines() []string {
	lines := []string{}
	input.Lines(func(s *bufio.Scanner) {lines = append(lines, s.Text())})
	return lines
}

func main() {
	fmt.Println()
	part1 := 0
	seq2totalScore := map[[4]int]int{}
	for _, line := range Lines() {
		n, _ := strconv.Atoi(line)
		prices := []int{}
		for i := 0; i < 2000; i++ {
			n = newSecret(n)
			prices = append(prices, n%10)
		}
		part1 += n
		changes := []int{}
		for i := 0; i < len(prices)-1; i++ {
			changes = append(changes, prices[i+1]-prices[i])
		}
		seq2score := map[[4]int]int{}
		for i := 0; i < len(changes)-3; i++ {
			seq := [4]int{changes[i], changes[i+1], changes[i+2], changes[i+3]}
			_, ok := seq2score[seq]
			if !ok {
				seq2score[seq] = prices[i+4]
			}
		}
		for seq, score := range seq2score {
			_, ok := seq2totalScore[seq]
			if ok {
				seq2totalScore[seq] += score
			} else {
				seq2totalScore[seq] = score
			}
		}
	}
	fmt.Println(part1)

	max := 0
	for s := range maps.Values(seq2totalScore) {
		if s > max {
			max = s
		}
	}
	fmt.Println(max)
}
