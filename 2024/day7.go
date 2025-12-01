package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7_1() {
	lines := strings.Split(strings.TrimSpace(string(ReadInput(7, false))), "\n")
	sum := 0
	isOk := func(total int, nums []int) bool {
		return tryOperations(nums[0], nums[1:], total)
	}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		total, _ := strconv.Atoi(parts[0])

		numStrs := strings.Fields(parts[1])
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			nums[i], _ = strconv.Atoi(numStr)
		}

		// check if possible to reach total
		if isOk(total, nums) {
			sum += total
		}
	}

	fmt.Printf("Day7_1: %d\n", sum)
}

// recursively tries all possible combinations of + and * operations
func tryOperations(current int, rest []int, target int) bool {
	if len(rest) == 0 {
		return current == target
	}

	// try addition
	if tryOperations(current+rest[0], rest[1:], target) {
		return true
	}

	// try multiplication
	if tryOperations(current*rest[0], rest[1:], target) {
		return true
	}

	return false
}

func Day7_2() {
	lines := strings.Split(strings.TrimSpace(string(ReadInput(7, false))), "\n")
	sum := 0

	isOk := func(total int, nums []int) bool {
		return tryOperationsP2(nums[0], nums[1:], total)
	}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))

		numStrs := strings.Fields(strings.TrimSpace(parts[1]))
		nums := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		if isOk(target, nums) {
			sum += target
		}
	}

	fmt.Printf("Day7_2: %d\n", sum)
}

func tryOperationsP2(current int, rest []int, target int) bool {
	if len(rest) == 0 {
		return current == target
	}

	if tryOperationsP2(current+rest[0], rest[1:], target) {
		return true
	}

	if tryOperationsP2(current*rest[0], rest[1:], target) {
		return true
	}

	concatenateNumbers := func(a, b int) int {
		bStr := strconv.Itoa(b)
		aStr := strconv.Itoa(a)
		result, _ := strconv.Atoi(aStr + bStr)
		return result
	}

	concatenated := concatenateNumbers(current, rest[0])
	if tryOperationsP2(concatenated, rest[1:], target) {
		return true
	}

	return false
}
