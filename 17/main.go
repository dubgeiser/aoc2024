package main

import (
	"aoc/lib/slice"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	A, B, C int
	ins     []int
}

const ADV, BXL, BST, JNZ, BXC, OUT, BDV, CDV = 0, 1, 2, 3, 4, 5, 6, 7

func (p *Program) Run() string {
	var output []int
	i := 0
	for i < len(p.ins) {
		opcode := p.ins[i]
		operand := p.ins[i+1]
		if opcode == ADV {
			// From https://stackoverflow.com/questions/5801008/go-and-operators
			//
			//  So n << x is "n times 2, x times".
			//  And y >> z is "y divided by 2, z times.
			//
			// => p.A / 2^combo(operand)
			p.A = p.A >> p.combo(operand)
		} else if opcode == BXL {
			p.B = p.B ^ operand

		} else if opcode == BST {
			p.B = p.combo(operand) % 8
		} else if opcode == JNZ {
			if p.A != 0 {
				i = operand
				continue
			}
		} else if opcode == BXC {
			p.B = p.B ^ p.C
		} else if opcode == OUT {
			output = append(output, p.combo(operand)%8)
		} else if opcode == BDV {
			p.B = p.A >> p.combo(operand)
		} else if opcode == CDV {
			p.C = p.A >> p.combo(operand)
		}
		i += 2
	}
	return strings.Join(slice.Map(strconv.Itoa, output), ",")
}

func (p *Program) combo(operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	if operand == 4 {
		return p.A
	}
	if operand == 5 {
		return p.B
	}
	if operand == 6 {
		return p.C
	}
	panic("UNKNOWN OPERAND")
}

func main() {
	fmt.Println()
	content, err := os.ReadFile("./input")
	if err != nil {
		panic("Cannot read file")
	}

	re := regexp.MustCompile(`\d+`)
	nums := slice.Map(slice.Int, re.FindAllString(string(content), -1))
	prg := &Program{A: nums[0], B: nums[1], C: nums[2], ins: nums[3:]}
	fmt.Println(prg.Run())
}
