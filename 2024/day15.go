package main

import (
	"fmt"
	"strings"
)

type Point15 struct {
	x, y int
}

func Day15_1() {
	input := string(ReadInput(15, true))
	parts := strings.Split(input, "\n\n")
	grid := strings.Split(parts[0], "\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")

	var robot Point15
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '@' {
				robot = Point15{x, y}
				break
			}
		}
	}
	warehouse := make([][]rune, len(grid))
	for i := range grid {
		warehouse[i] = []rune(grid[i])
	}

	for _, move := range moves {
		dx, dy := 0, 0
		switch move {
		case '^':
			dy = -1
		case 'v':
			dy = 1
		case '<':
			dx = -1
		case '>':
			dx = 1
		}

		newX, newY := robot.x+dx, robot.y+dy

		if warehouse[newY][newX] == '#' {
			continue
		}

		if warehouse[newY][newX] == 'O' {
			nextX, nextY := newX+dx, newY+dy
			// Check if next position is valid and empty
			if nextY < 0 || nextY >= len(warehouse) || nextX < 0 || nextX >= len(warehouse[0]) ||
				warehouse[nextY][nextX] == '#' || warehouse[nextY][nextX] == 'O' {
				continue // Can't push the box, skip this move
			}

			// Move the box
			warehouse[nextY][nextX] = 'O'
			// Move the robot
			warehouse[newY][newX] = '@'
			warehouse[robot.y][robot.x] = '.'
			robot.x, robot.y = newX, newY
		} else if warehouse[newY][newX] == '.' {
			// Move the robot to empty space
			warehouse[newY][newX] = '@'
			warehouse[robot.y][robot.x] = '.'
			robot.x, robot.y = newX, newY
		}
	}

	sum := 0
	for y := range warehouse {
		for x := range warehouse[y] {
			if warehouse[y][x] == 'O' {
				sum += y*100 + x
			}
		}
	}

	fmt.Println(sum)
}
