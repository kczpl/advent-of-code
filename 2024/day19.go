package main

import (
	"fmt"
	"os"
	"strings"
)

type memoKey struct {
	design string
}

func valid(patterns []string, design string, memo map[memoKey]bool) bool {
	key := memoKey{design}

	if result, exists := memo[key]; exists {
		return result
	}

	if len(design) == 0 {
		return true
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining := design[len(pattern):]
			if valid(patterns, remaining, memo) {
				memo[key] = true
				return true
			}
		}
	}

	memo[key] = false
	return false
}

func countWays(patterns []string, design string, memo map[memoKey]int) int {
	key := memoKey{design}

	if count, exists := memo[key]; exists {
		return count
	}

	if len(design) == 0 {
		return 1
	}

	total := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining := design[len(pattern):]
			total += countWays(patterns, remaining, memo)
		}
	}

	memo[key] = total
	return total
}

func Day19() {
	data, err := os.ReadFile("data/input19.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(data), "\n\n")
	if len(parts) != 2 {
		panic("Invalid input format")
	}

	patternLine := strings.TrimSpace(parts[0])
	patterns := strings.Split(patternLine, ",")
	for i := range patterns {
		patterns[i] = strings.TrimSpace(patterns[i])
	}

	designs := strings.Split(strings.TrimSpace(parts[1]), "\n")

	validMemo := make(map[memoKey]bool)
	validCount := 0
	for _, design := range designs {
		if valid(patterns, design, validMemo) {
			validCount++
		}
	}
	fmt.Println(validCount)

	waysMemo := make(map[memoKey]int)
	totalWays := 0
	for _, design := range designs {
		ways := countWays(patterns, design, waysMemo)
		totalWays += ways
	}
	fmt.Println(totalWays)
}