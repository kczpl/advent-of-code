package main

import (
	"fmt"
	"strings"
)

func Day10() {
	i := strings.Split(string(ReadInput(10, false)), "\n")
	m := make([][]int, len(i))

	var trailheads [][2]int

	for x, line := range i {
		m[x] = make([]int, len(line))
		for y, ch := range line {
			m[x][y] = int(ch - '0')
			if m[x][y] == 0 {
				trailheads = append(trailheads, [2]int{x, y})
			}
		}
	}

	sum1 := 0
	for _, th := range trailheads {
		score := bfs(m, th)
		sum1 += score
	}

	fmt.Println(sum1)

	sum2 := 0
	for _, th := range trailheads {
		score := bfs2(m, th)
		sum2 += score
	}

	fmt.Println(sum2)
}

func bfs(m [][]int, start [2]int) int {
	// keep track of visited
	visited := make(map[[2]int]bool)

	// longest path is nine
	reachableNine := make(map[[2]int]bool)

	// queue starts with start trailhead
	queue := [][2]int{start}

	visited[start] = true

	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	rows, cols := len(m), len(m[0])

	// bfs loop
	for len(queue) > 0 {
		// get next position from queue
		current := queue[0]
		queue = queue[1:]

		if m[current[0]][current[1]] == 9 {
			reachableNine[current] = true
			continue
		}

		// check neighboring positions
		for _, dir := range dirs {
			next := [2]int{current[0] + dir[0], current[1] + dir[1]}

			// check if position is within grid bounds
			if next[0] < 0 || next[0] >= rows || next[1] < 0 || next[1] >= cols {
				continue
			}

			// check if position has been visited
			if visited[next] {
				continue
			}
			// check if height increases by exactly 1
			if m[next[0]][next[1]] != m[current[0]][current[1]]+1 {
				continue
			}

			// add to queue once
			visited[next] = true
			queue = append(queue, next)
		}
	}
	return len(reachableNine)
}

func bfs2(m [][]int, start [2]int) int {
	// keep track of visited paths to each position
	paths := make(map[[2]int]int)

	// queue starts with start trailhead
	queue := [][2]int{start}
	paths[start] = 1

	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	rows, cols := len(m), len(m[0])

	// total paths to all 9s
	totalPaths := 0

	// bfs loop
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// if we reach a 9, add the number of paths to this 9 to total
		if m[current[0]][current[1]] == 9 {
			totalPaths += paths[current]
			continue
		}

		// check neighboring positions
		for _, dir := range dirs {
			next := [2]int{current[0] + dir[0], current[1] + dir[1]}

			// check if position is within grid bounds
			if next[0] < 0 || next[0] >= rows || next[1] < 0 || next[1] >= cols {
				continue
			}

			// check if height increases by exactly 1
			if m[next[0]][next[1]] != m[current[0]][current[1]]+1 {
				continue
			}

			// add paths to this position
			paths[next] += paths[current]

			// only add to queue if we haven't processed this position before
			if paths[next] == paths[current] {
				queue = append(queue, next)
			}
		}
	}
	return totalPaths
}
