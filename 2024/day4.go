package main

import (
	"fmt"
	"strings"
)

func Day4_1() {
	rawInput := ReadInput(4, false)
	input := strings.Split(string(rawInput), "\n")
	m := make([][]string, len(input))

	for i, line := range input {
		elements := strings.Split(line, "")
		m[i] = elements
	}

	directions := [][2]int{
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
	}

	rows := len(m)
	cols := len(m[0])
	sum := 0

	findMatchinDir := func(m [][]string, startX, startY, dx, dy, rows, cols int) bool {
		target := "XMAS"

		// bounds
		endX := startX + dx*3
		endY := startY + dy*3

		// out if matrix bounds check
		if endX < 0 || endX >= rows || endY < 0 || endY >= cols {
			return false
		}

		for i := 0; i < len(target); i++ {
			x := startX + dx*i
			y := startY + dy*i

			if string(m[x][y]) != string(target[i]) {
				return false
			}
		}

		return true
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				if findMatchinDir(m, x, y, dir[0], dir[1], len(m), len(m[0])) {
					sum++
				}
			}
		}
	}
	fmt.Printf("Day4_1: %d\n", sum)
}

func Day4_2() {
	rawInput := ReadInput(4, false)
	input := strings.Split(string(rawInput), "\n")
	m := make([][]string, len(input))

	for i, line := range input {
		elements := strings.Split(line, "")
		m[i] = elements
	}

	// not needed here but still helps in visualizing fiagonals
	// directions := [][2]int{
	// 	{1, 1},   // down-right
	// 	{1, -1},  // down-left
	// 	{-1, -1}, // up-left
	// 	{-1, 1},  // up-right
	// }

	rows := len(m)
	cols := len(m[0])
	sum := 0

	findXMatch := func(m [][]string, startX, startY, rows, cols int) bool {
		if startX+1 >= rows || startX-1 < 0 || startY+1 >= cols || startY-1 < 0 {
			return false
		}

		// diagonal /
		diag1 := m[startX-1][startY+1] == "M" && m[startX][startY] == "A" && m[startX+1][startY-1] == "S" ||
			m[startX-1][startY+1] == "S" && m[startX][startY] == "A" && m[startX+1][startY-1] == "M"

		// diagonal \
		diag2 := m[startX-1][startY-1] == "M" && m[startX][startY] == "A" && m[startX+1][startY+1] == "S" ||
			m[startX-1][startY-1] == "S" && m[startX][startY] == "A" && m[startX+1][startY+1] == "M"

		return diag1 && diag2
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if m[x][y] == "A" {
				if findXMatch(m, x, y, rows, cols) {
					sum++
				}
			}
		}
	}
	fmt.Printf("Day4_2: %d\n", sum)
}
