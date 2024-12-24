package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() string {
	if len(os.Args) <= 1 {
		return "input"
	}
	return os.Args[1]
}

type LineProcessor interface {
	ProcessLine(i int, line string)
}

func ReadLines(fn string, lp LineProcessor) (int, error) {
	file, err := os.Open(fn)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var i int
	for i = 0; scanner.Scan(); i++ {
		lp.ProcessLine(i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return i, nil
}

func ReadTwoParts(fn string) (string, string) {
	bContent, err := os.ReadFile(fn)
	if err != nil {
		panic("Cannot read file!")
	}
	parts := strings.Split(string(bContent), "\n\n")
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func Read() string {
	fn := Input()
	content, err := os.ReadFile(fn)
	if err != nil {
		panic(fmt.Sprintf("Canot read [%s]", fn))
	}
	return strings.TrimSpace(string(content))
}
