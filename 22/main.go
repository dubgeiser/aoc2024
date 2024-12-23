package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	raw, err := os.ReadFile("./input")
	if err != nil {
		panic("Cannot read input!")
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	count := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			n = (n ^ (n * 64)) % 16777216
			n = (n ^ (n / 32)) % 16777216
			n = (n ^ (n * 2048)) % 16777216
		}
		count += n
	}
	fmt.Println(count)
}
