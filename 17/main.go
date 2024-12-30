package main

import (
	"aoc/lib/input"
	"aoc/lib/slice"
	"fmt"
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
	content := input.Blob()

	re := regexp.MustCompile(`\d+`)
	nums := slice.Map(slice.Int, re.FindAllString(content, -1))

	// --[ Part 1 ]--
	prg := &Program{A: nums[0], B: nums[1], C: nums[2], ins: nums[3:]}
	fmt.Println(prg.Run())

	// --[ Part 2 ]--
	fmt.Println(FindQuine(nums[3:], 0))
}

// Figuring out what the program does
//
// * Sample: 0,3,5,4,3,0
// 0. (0,3) ADV A>>3
// 1. (5,4) OUT A%8
// 2. (3,0) JNZ if A != 0 goto 0
//
// This is the equivalent of:   // See ExecuteSampleProgram
// output = []
// for
//
//	A = A>>3
//	output = append(output, A%8)
//	if A == 0:
//	    break
//
// * Input: 2,4,1,5,7,5,1,6,4,2,5,5,0,3,3,0
// 0. (2,4) BST B = A%8
// 1. (1,5) BXL B = B^5
// 2. (7,5) CDV C = A>>B
// 3. (1,6) BXL B = B^6
// 4. (4,2) BXC B = B^C
// 5. (5,5) OUT B%8
// 6. (0,3) ADV A = A>>3
// 7. (3,0) JNZ if A!=0 goto 0
//
// This is the equivalent of a function: f(A)
//
//		output = []
//		for
//		    B = A%8
//		    B = B^5
//		    C = A>>B
//		    B = B^6
//		    B = B^C
//		    output = append(output, B%8)
//		    A = A>>3
//		    if A==0: break
//	 return output
//
// If f(A) is a quine => A = 0 when JNZ is reached after the last OUT.
// If not, the loop would run again, and thus _not_ be a quine.
//
//           last loop: A>>3 == 0 (true for A=0..7)
// Second to last loop: A>>3 == 1 (true for A=8..15)
// ...                  A>>3 == 2 (true for A=16..23)
// ...					A>>3 == 3 (true for A=24..31)
// ...                  ...
// ...                  A>>3 == n (true for A=(A<<3)+0 ..(A<<3)+7
//
// So we can check different values of A in a loop
// Something like:
// for i in (0,7):
//	A, B, C = A<<3 + i, 0, 0
//
// Recursion, it's in there...
// We can write this as a function that operates on (part of) a program:
// We exhaust the program from R->L and check if our OUT is equal to the last
// digit in the program.
func FindQuine(prg []int, a int) int {
	// We've checked every element in our program
	if len(prg) == 0 {
		return a
	}
	for i := 0; i < 8; i++ {
		A := (a<<3) + i
		B := A % 8
		B = B ^ 5
		C := A >> B
		B = B ^ 6
		B = B ^ C
		out := B % 8
		// Are we in sync with the last element of our program?
		// Recurse to find A for the previous digit of our program.
		if out == prg[len(prg)-1] {
			q := FindQuine(prg[:len(prg)-1], A)
			if q == -1 {
				continue
			}
			return q
		}
	}
	return -1
}

func ExecuteProgram(A, B, C int) []int {
	output := []int{}
	for {
		B = A % 8
		B = B ^ 5
		C = A >> B
		B = B ^ 6
		B = B ^ C
		output = append(output, B%8)
		A = A >> 3
		if A == 0 {
			break
		}
	}
	return output
}

func ExecuteSampleProgram(A int) []int {
	output := []int{}
	fmt.Println("A: ", A)
	for {
		A = A >> 3
		output = append(output, A%8)
		fmt.Println("A: ", A)
		if A == 0 {
			break
		}
	}
	return output
}
