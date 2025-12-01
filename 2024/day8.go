package main

import (
	"fmt"
	"math"
	"strings"
)

func Day8_1() {
	rawInput := ReadInput(8, false)
	input := strings.Split(string(rawInput), "\n")
	m := make([][]string, len(input))

	for i, line := range input {
		elements := strings.Split(line, "")
		m[i] = elements
	}

	antinodesPositions := make(map[string]bool)

	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] == "." {
				continue
			}

			freq := m[x][y] // a

			for x2 := 0; x2 < len(m); x2++ {
				for y2 := 0; y2 < len(m[x2]); y2++ {
					// same field
					if x == x2 && y == y2 {
						continue
					}

					// same freq found
					// p1 = (4,3) p2 = (5,5)
					// a1 = (3,1) a2 = (6,7)
					// ax1 = 2*4 - 5 = 3
					// ay1 = 2*3 - 5 = 1
					if m[x2][y2] == freq {
						// first antinode
						ax1 := 2*x - x2
						ay1 := 2*y - y2

						if ax1 >= 0 && ax1 < len(m) && ay1 >= 0 && ay1 < len(m[0]) {
							antinodesPositions[fmt.Sprintf("%d,%d", ax1, ay1)] = true
						}

						// second antinode
						ax2 := 2*x2 - x
						ay2 := 2*y2 - y

						if ax2 >= 0 && ax2 < len(m) && ay2 >= 0 && ay2 < len(m[0]) {
							antinodesPositions[fmt.Sprintf("%d,%d", ax2, ay2)] = true
						}
					}
				}
			}
		}
	}

	fmt.Println(len(antinodesPositions))
}

func Day8_2() {
	rawInput := ReadInput(8, false)
	input := strings.Split(string(rawInput), "\n")
	m := make([][]string, len(input))

	for i, line := range input {
		elements := strings.Split(line, "")
		m[i] = elements
	}

	antinodesPositions := make(map[string]bool)

	reduce := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// for each antenna
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if m[x][y] == "." {
				continue
			}

			freq := m[x][y]

			for x2 := 0; x2 < len(m); x2++ {
				for y2 := 0; y2 < len(m[0]); y2++ {
					if x == x2 && y == y2 {
						continue
					}

					if m[x2][y2] == freq {
						// line function between two antennas
						// p1 = (4,3) p2 = (5,5)
						// dx = 1 dy = -2
						dx := x2 - x
						dy := y2 - y

						// reduce it to find every possible position on the line
						reducedDir := reduce(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
						if reducedDir > 0 {
							dx /= reducedDir
							dy /= reducedDir
						}

						currentX, currentY := x, y

						// forward
						for currentX >= 0 && currentX < len(m) && currentY >= 0 && currentY < len(m[0]) {
							antinodesPositions[fmt.Sprintf("%d,%d", currentX, currentY)] = true
							currentX += dx
							currentY += dy
						}

						// backward
						currentX, currentY = x-dx, y-dy
						for currentX >= 0 && currentX < len(m) && currentY >= 0 && currentY < len(m[0]) {
							antinodesPositions[fmt.Sprintf("%d,%d", currentX, currentY)] = true
							currentX -= dx
							currentY -= dy
						}
					}
				}
			}
		}
	}

	fmt.Println(len(antinodesPositions))
}
