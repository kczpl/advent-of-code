package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Registers struct {
	A, B, C int
}

func getComboVal(op int, regs Registers) int {
	switch {
	case op < 4:
		return op
	case op == 4:
		return regs.A
	case op == 5:
		return regs.B
	case op == 6:
		return regs.C
	default:
		return 0
	}
}

func solve(regs Registers, program []int) []int {
	outputs := []int{}
	ip := 0

	for ip < len(program) {
		op := program[ip]
		operand := 0
		if ip+1 < len(program) {
			operand = program[ip+1]
		}

		switch op {
		case 0: // adv
			regs.A /= 1 << getComboVal(operand, regs)
		case 1: // bxl
			regs.B ^= operand
		case 2: // bst
			regs.B = getComboVal(operand, regs) % 8
		case 3: // jnz
			if regs.A != 0 {
				ip = operand
				continue
			}
		case 4: // bxc
			regs.B ^= regs.C
		case 5: // out
			outputs = append(outputs, getComboVal(operand, regs)%8)
		case 6: // bdv
			regs.B = regs.A / (1 << getComboVal(operand, regs))
		case 7: // cdv
			regs.C = regs.A / (1 << getComboVal(operand, regs))
		}
		ip += 2
	}
	return outputs
}

func Day17_1() {
	lines := strings.Split(strings.TrimSpace(string(ReadInput(17, false))), "\n")

	regs := Registers{}
	regs.A, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	regs.B, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])
	regs.C, _ = strconv.Atoi(strings.Split(lines[2], ": ")[1])

	programStr := strings.Split(lines[len(lines)-1], ": ")[1]
	programNums := strings.Split(programStr, ",")
	program := make([]int, len(programNums))
	for i, num := range programNums {
		program[i], _ = strconv.Atoi(num)
	}

	// Run program and join outputs
	outputs := solve(regs, program)
	result := make([]string, len(outputs))
	for i, v := range outputs {
		result[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(result, ","))
}
