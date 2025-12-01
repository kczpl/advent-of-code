package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day11_1() {
	line := strings.Split(string(ReadInput(11, false)), " ")
	stones := make([]int, 0)

	for _, stone := range line {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, num)
	}

	buffor := make([]int, 0)

	for i := 0; i < 25; i++ {
		fmt.Println(i)
		for _, stone := range stones {
			switch {
			case stone == 0:
				buffor = append(buffor, 1)
			case len(strconv.Itoa(stone))%2 == 0:
				strStone := strconv.Itoa(stone)
				mid := len(strStone) / 2
				left, _ := strconv.Atoi(strStone[:mid])
				right, _ := strconv.Atoi(strStone[mid:])
				buffor = append(buffor, left, right)
			default:
				buffor = append(buffor, stone*2024)
			}
		}

		stones = make([]int, len(buffor))
		copy(stones, buffor)
		buffor = make([]int, 0)
	}

	fmt.Println(len(stones))
}

func Day11_2() {
	line := strings.Split(string(ReadInput(11, false)), " ")
	stones := make([]string, 0)
	for _, stone := range line {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, strconv.Itoa(num))
	}

	// each stone is a key and value is the number how many times it appears
	oldStones := make(map[string]int)
	for _, stone := range stones {
		oldStones[stone]++
	}

	for i := 0; i < 75; i++ {
		newStones := make(map[string]int)
		for stone, count := range oldStones {

			if stone == "0" {
				newStones["1"] += count
			} else if len(stone)%2 == 0 {
				mid := len(stone) / 2
				left, _ := strconv.Atoi(stone[:mid])
				right, _ := strconv.Atoi(stone[mid:])
				newStones[strconv.Itoa(left)] += count
				newStones[strconv.Itoa(right)] += count
			} else {
				num, _ := strconv.Atoi(stone)
				newStones[strconv.Itoa(num*2024)] += count
			}

		}
		oldStones = newStones
	}

	sum := 0
	for _, count := range oldStones {
		sum += count
	}
	fmt.Println(sum)
}
