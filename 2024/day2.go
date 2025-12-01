package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2_1() {
	input := ReadInput(2, false)
	s := string(input)
	res := 0

	for _, row := range strings.Split(s, "\n") {
		reports := strings.Fields(row)

		isValid := true
		var goesUp bool

		for idx := 0; idx < len(reports)-1; idx++ {
			number, _ := strconv.Atoi(reports[idx])
			next_number, _ := strconv.Atoi(reports[idx+1])

			if next_number == number {
				isValid = false
				break
			}

			if idx == 0 {
				goesUp = next_number > number
			} else {
				if (next_number > number) != goesUp {
					isValid = false
					break
				}
			}

			if goesUp {
				diff := next_number - number
				if diff < 1 || diff > 3 {
					isValid = false
					break
				}
			} else {
				diff := number - next_number
				if diff < 1 || diff > 3 {
					isValid = false
					break
				}
			}
		}

		if isValid {
			res++
		}
	}
	fmt.Printf("Day2_1: %d\n", res)
}

func Day2_2() {
	input := ReadInput(2, false)
	s := string(input)
	res := 0

	checkSequence := func(reports []string) bool {
		number, _ := strconv.Atoi(reports[0])
		next_number, _ := strconv.Atoi(reports[1])
		if number == next_number {
			return false
		}
		goesUp := next_number > number

		diff := 0
		if goesUp {
			diff = next_number - number
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			diff = number - next_number
			if diff < 1 || diff > 3 {
				return false
			}
		}

		for i := 1; i < len(reports)-1; i++ {
			current, _ := strconv.Atoi(reports[i])
			next, _ := strconv.Atoi(reports[i+1])

			if current == next {
				return false
			}

			if goesUp {
				diff = next - current
				if diff < 1 || diff > 3 {
					return false
				}
			} else {
				diff = current - next
				if diff < 1 || diff > 3 {
					return false
				}
			}
		}
		return true
	}

	for _, row := range strings.Split(s, "\n") {
		reports := strings.Fields(row)
		isValid := checkSequence(reports)

		if !isValid {
			for i := 0; i < len(reports); i++ {
				newReports := append([]string{}, reports[:i]...)
				newReports = append(newReports, reports[i+1:]...)

				if checkSequence(newReports) {
					isValid = true
					break
				}
			}
		}

		if isValid {
			res++
		}
	}

	fmt.Printf("Day2_2: %d\n", res)
}
