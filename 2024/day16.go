package main

import (
	"fmt"
	"sort"
	"strings"
)

type Point16 struct {
	x, y int
}

type State16 struct {
	position  Point16
	direction int // 0=right, 1=down, 2=left, 3=up
	score     int
}

func Day16_1() {
	input := strings.Split(string(ReadInput(16, false)), "\n")
	grid := make([][]rune, len(input))
	for i := range input {
		grid[i] = []rune(input[i])
	}

	var start, end Point16
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'S' {
				start = Point16{x, y}
			}
			if grid[y][x] == 'E' {
				end = Point16{x, y}
			}
		}
	}

	result := findPathP1(grid, start, end)
	fmt.Println(result) // 85420
}

func findPathP1(grid [][]rune, start, end Point16) int {
	queue := make([]State16, 0)
	visited := make(map[string]int)

	queue = append(queue, State16{start, 0, 0}) //right

	move := func(p Point16, direction int) Point16 {
		switch direction {
		case 0: // right
			return Point16{p.x + 1, p.y}
		case 1: // down
			return Point16{p.x, p.y + 1}
		case 2: // left
			return Point16{p.x - 1, p.y}
		case 3: // up
			return Point16{p.x, p.y - 1}
		}
		return p
	}

	isValidPosition := func(p Point16, grid [][]rune) bool {
		if p.y < 0 || p.y >= len(grid) || p.x < 0 || p.x >= len(grid[0]) {
			return false
		}
		return grid[p.y][p.x] != '#'
	}

	stateKey := func(p Point16, direction int) string {
		return fmt.Sprintf("%d,%d,%d", p.x, p.y, direction)
	}

	sortQueue := func(queue []State16) {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].score < queue[j].score
		})
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position == end {
			return current.score
		}

		newPosition := move(current.position, current.direction)

		if isValidPosition(newPosition, grid) {
			newState := State16{newPosition, current.direction, current.score + 1}
			key := stateKey(newPosition, current.direction)

			if score, exists := visited[key]; !exists || score > newState.score {
				visited[key] = newState.score
				queue = append(queue, newState)
			}
		}

		// try turn left and right
		for _, newDir := range []int{(current.direction + 1) % 4, (current.direction + 3) % 4} {
			newState := State16{current.position, newDir, current.score + 1000}
			key := stateKey(current.position, newDir)

			if score, exists := visited[key]; !exists || score > newState.score {
				visited[key] = newState.score
				queue = append(queue, newState)
			}
		}

		sortQueue(queue)
	}
	return -1
}
