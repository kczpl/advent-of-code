package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Point14 struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}

type Robot struct {
	Position Point14
	Velocity Velocity
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func Day14_1() {
	test := false
	input := string(ReadInput(14, test))
	lines := strings.Split(input, "\n")
	seconds := 100

	gridWidth := 101
	gridHeight := 103
	if test {
		gridWidth = 11
		gridHeight = 7
	}

	grid := [][]int{}
	for i := 0; i < gridHeight; i++ {
		grid = append(grid, make([]int, gridWidth))
	}

	robots := []Robot{}
	for _, line := range lines {
		startPoint := regexp.MustCompile(`p=(-?\d+),(-?\d+)`).FindStringSubmatch(line)
		velocity := regexp.MustCompile(`v=(-?\d+),(-?\d+)`).FindStringSubmatch(line)
		robot := Robot{
			Position: Point14{X: parseInt(startPoint[1]), Y: parseInt(startPoint[2])},
			Velocity: Velocity{X: parseInt(velocity[1]), Y: parseInt(velocity[2])},
		}
		robots = append(robots, robot)
	}

	// X positive is right, Y positive is down

	// grid width is 11 and height is 7
	// x=4, y=1 ; next second x=6, y=5 ; velocity x=2, y=-3; positionY = -2 ends up in 5; so
	// ((x % n) + n) % n

	// it's called modulo wrap-around or toroidal wrapping TIL

	for i := 0; i < seconds; i++ {
		for ridx := range robots {
			robots[ridx].Position.X += robots[ridx].Velocity.X
			robots[ridx].Position.Y += robots[ridx].Velocity.Y

			robots[ridx].Position.X = ((robots[ridx].Position.X % gridWidth) + gridWidth) % gridWidth
			robots[ridx].Position.Y = ((robots[ridx].Position.Y % gridHeight) + gridHeight) % gridHeight
		}
	}

	// robots can have same positions
	for _, robot := range robots {
		grid[robot.Position.Y][robot.Position.X]++
	}

	// quadrants
	midX := gridWidth / 2
	midY := gridHeight / 2
	q1, q2, q3, q4 := 0, 0, 0, 0

	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			// middle
			if x == midX || y == midY {
				continue
			}

			count := grid[y][x]
			if x < midX && y < midY {
				q1 += count
			} else if x > midX && y < midY {
				q2 += count
			} else if x < midX && y > midY {
				q3 += count
			} else if x > midX && y > midY {
				q4 += count
			}
		}
	}
	result := q1 * q2 * q3 * q4
	fmt.Println(result)
}
