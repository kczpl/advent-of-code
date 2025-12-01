package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day6_1() {
	data := string(ReadInput(6, false))
	lines := strings.Split(data, "\n")
	matrix := make([][]string, 0, len(lines))
	startCoords := []int{0, 0}

	for _, line := range lines {
		elements := strings.Split(line, "")
		matrix = append(matrix, elements)
	}

	for rIdx, row := range matrix {
		if slices.Contains(row, "^") {
			startCoords[0] = rIdx
			startCoords[1] = slices.Index(row, "^")
			break
		}
	}

	visitedPositions := make(map[[2]int]bool)

	directions := [][2]int{
		{0, 1},  // right 0
		{1, 0},  // down 1
		{0, -1}, // left 2
		{-1, 0}, // up 3
	}

	currentDirection := directions[3]

	checkDirToRight := func(currentDirection [2]int) [2]int {
		if currentDirection == directions[0] {
			return directions[1]
		} else if currentDirection == directions[1] {
			return directions[2]
		} else if currentDirection == directions[2] {
			return directions[3]
		}
		return directions[0]
	}

	for {
		nextStep := []int{startCoords[0] + currentDirection[0], startCoords[1] + currentDirection[1]}

		if matrix[nextStep[0]][nextStep[1]] == "#" {
			currentDirection = checkDirToRight(currentDirection)
		} else {
			startCoords = nextStep
			visitedPositions[[2]int{startCoords[0], startCoords[1]}] = true
		}

		if startCoords[0] == 0 || startCoords[0] == len(matrix)-1 || startCoords[1] == 0 || startCoords[1] == len(matrix[0])-1 {
			break
		}
	}

	fmt.Printf("Day6_1: %d\n", len(visitedPositions))
}

func Day6_2() {
	data := string(ReadInput(6, false))
	lines := strings.Split(data, "\n")
	matrix := make([][]string, 0, len(lines))
	startCoords := []int{0, 0}

	for _, line := range lines {
		elements := strings.Split(line, "")
		matrix = append(matrix, elements)
	}

	for rIdx, row := range matrix {
		if slices.Contains(row, "^") {
			startCoords[0] = rIdx
			startCoords[1] = slices.Index(row, "^")
			break
		}
	}

	foundBlockers := 0

	directions := [][2]int{
		{0, 1},  // right 0
		{1, 0},  // down 1
		{0, -1}, // left 2
		{-1, 0}, // up 3
	}

	checkDirToRight := func(currentDirection [2]int) [2]int {
		if currentDirection == directions[0] {
			return directions[1]
		} else if currentDirection == directions[1] {
			return directions[2]
		} else if currentDirection == directions[2] {
			return directions[3]
		}
		return directions[0]
	}

	looping := func(matrix [][]string, startCoords []int) {
		currentDirection := directions[3]
		visitedPositions := make(map[[2]int]int)

		for {
			nextStep := []int{startCoords[0] + currentDirection[0], startCoords[1] + currentDirection[1]}

			if matrix[nextStep[0]][nextStep[1]] == "#" {
				currentDirection = checkDirToRight(currentDirection)
			} else {
				startCoords = nextStep
				visitedPositions[[2]int{startCoords[0], startCoords[1]}] += 1
				if startCoords[0] == 0 || startCoords[0] == len(matrix)-1 || startCoords[1] == 0 || startCoords[1] == len(matrix[0])-1 {
					break
				}
				continue
			}

			for _, count := range visitedPositions {
				if count >= 5 { // 9 works
					foundBlockers++
					return
				}
			}
		}
	}

	for rIdx, row := range matrix {
		for cIdx := range row {
			// fmt.Println(rIdx, cIdx)
			if matrix[rIdx][cIdx] == "#" {
				continue
			}

			if matrix[rIdx][cIdx] == "^" {
				continue
			}

			newMatrix := make([][]string, len(matrix))

			for i := range matrix {
				newMatrix[i] = make([]string, len(matrix[i]))
				copy(newMatrix[i], matrix[i])
			}

			newStartCoords := []int{startCoords[0], startCoords[1]}
			newMatrix[rIdx][cIdx] = "#"

			looping(newMatrix, newStartCoords)
		}
	}

	fmt.Printf("Day6_2: %d\n", foundBlockers)
}
