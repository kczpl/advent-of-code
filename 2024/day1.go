package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1_1() {
	input := ReadInput(1, false)
	s := string(input)
	pairs := strings.Split(s, "\n")

	var left, right, res []int

	for _, p := range pairs {
		pair := strings.Fields(p)
		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[len(pair)-1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		res = append(res, int(math.Abs(float64(left[i]-right[i]))))
	}

	sum := 0
	for _, num := range res {
		sum += num
	}

	fmt.Printf("Day1_1: %d\n", sum)
}

func Day1_2() {
	input := ReadInput(1, false)
	s := string(input)
	pairs := strings.Split(s, "\n")

	var left, right, res []int

	for _, p := range pairs {
		pair := strings.Fields(p)
		l, _ := strconv.Atoi(pair[0])
		r, _ := strconv.Atoi(pair[len(pair)-1])

		left = append(left, l)
		right = append(right, r)
	}

	for i := 0; i < len(left); i++ {
		occ := countInt(right, left[i])

		res = append(res, occ*left[i])
	}

	sum := 0
	for _, num := range res {
		sum += num
	}

	fmt.Printf("Day1_2: %d\n", sum)
}

func countInt(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}
