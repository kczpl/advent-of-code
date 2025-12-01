package main

import (
	"fmt"
	"strings"
)

// solved with help of Reddit and llm
func Day12_1() {
	input := strings.Split(string(ReadInput(12, false)), "\n")
	m := make([][]string, len(input))

	regions := make(map[string][][2]int)
	visited := make(map[string]bool)

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for x, line := range input {
		m[x] = make([]string, len(line))

		for y, ch := range line {
			m[x][y] = string(ch)
		}
	}

	var exploreRegion func(x, y int, char string) [2]int
	exploreRegion = func(x, y int, char string) [2]int {
		if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) ||
			m[x][y] != char || visited[fmt.Sprintf("%d,%d", x, y)] {
			return [2]int{0, 0}
		}

		visited[fmt.Sprintf("%d,%d", x, y)] = true
		result := [2]int{1, 0}

		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]
			if newX < 0 || newX >= len(m) || newY < 0 || newY >= len(m[0]) || m[newX][newY] != char {
				result[1]++
			}
		}

		for _, dir := range dirs {
			newX, newY := x+dir[0], y+dir[1]

			// recursive explore region
			subRegion := exploreRegion(newX, newY, char)
			result[0] += subRegion[0]
			result[1] += subRegion[1]
		}

		return result
	}

	sum := 0
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {

			if !visited[fmt.Sprintf("%d,%d", x, y)] {
				region := exploreRegion(x, y, m[x][y])

				if region[0] > 0 {
					regions[m[x][y]] = append(regions[m[x][y]], region)
				}
			}
		}
	}

	for _, regionList := range regions {
		for _, region := range regionList {
			sum += region[0] * region[1]
		}
	}
	fmt.Println(sum)
}
