package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3_1() {
	input := ReadInput(3, false)

	regex := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
	re := regexp.MustCompile(regex)

	matches := re.FindAllStringSubmatch(string(input), -1)

	sum := 0
	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		sum += (l * r)
	}

	fmt.Printf("Day3_1: %d\n", sum)
}

func Day3_2() {
	input := ReadInput(3, false)
	s := string(input)

	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(do\(\))|(don't\(\))`)
	matches := re.FindAllStringSubmatch(s, -1)

	enabled := true
	sum := 0

	for _, match := range matches {
		if match[1] != "" {
			if enabled {
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				sum += (num1 * num2)
			}
		} else if match[4] != "" {
			enabled = true
		} else if match[5] != "" {
			enabled = false
		}
	}
	fmt.Printf("Day3_2: %d\n", sum)
}
