package main

import (
	"fmt"
	"os"
)

func main() {
	// Day1_1()
	// Day1_2()
	// Day2_1()
	// Day2_2()
	// Day3_1()
	// Day3_2()
	// Day4_1()
	// Day4_2()
	// Day5_1()
	// Day5_2()
	// Day6_1()
	// Day6_2()
	// Day7_1()
	// Day7_2()
	// Day8_1()
	// Day8_2()
	// Day9()
	// Day10()
	// Day11_1()
	// Day11_2()
	// Day12_1()
	// Day13_1()
	// Day13_2()
	// Day14_1()
	// Day16_1()
	// Day17_1()
	Day19()
}

func ReadInput(day int, test bool) []byte {
	var data []byte

	if test {
		data, _ = os.ReadFile(fmt.Sprintf("data/sample%d.txt", day))
	} else {
		data, _ = os.ReadFile(fmt.Sprintf("data/input%d.txt", day))
	}
	return data
}
