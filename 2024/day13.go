package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Button struct {
	Point  Point
	DeltaX int
	DeltaY int
	Cost   int
}

type ButtonSet struct {
	Buttons []Button
	Target  Point
}

func Day13_1() {
	buttonSets := parseInput(string(ReadInput(13, false)))

	// numA(?) * deltaAX + numB(?) * deltaBX = targetX
	// numA(?) * deltaAY + numB(?) * deltaBY = targetY

	// numA = (targetX * deltaBY - targetY * deltaBX) / (deltaAX * deltaBY - deltaAY * deltaBX)
	// numB = (deltaAX * targetY - deltaAY * targetX) / (deltaAX * deltaBY - deltaAY * deltaBX)

	allButtonHits := []int{}
	for _, set := range buttonSets {
		costA := set.Buttons[0].Cost
		costB := set.Buttons[1].Cost
		deltaAX := set.Buttons[0].DeltaX
		deltaAY := set.Buttons[0].DeltaY
		deltaBX := set.Buttons[1].DeltaX
		deltaBY := set.Buttons[1].DeltaY
		targetX := set.Target.X
		targetY := set.Target.Y

		numA := float64(targetX*deltaBY-targetY*deltaBX) / float64(deltaAX*deltaBY-deltaAY*deltaBX)
		numB := float64(deltaAX*targetY-deltaAY*targetX) / float64(deltaAX*deltaBY-deltaAY*deltaBX)

		if numA == float64(int(numA)) && numB == float64(int(numB)) &&
			numA >= 0 && numB >= 0 && numA <= 100 && numB <= 100 {
			hitsCost := int(numA)*costA + int(numB)*costB
			allButtonHits = append(allButtonHits, hitsCost)
		}
	}

	totalCost := 0
	for _, cost := range allButtonHits {
		totalCost += cost
	}

	fmt.Println(totalCost)
}

func Day13_2() {
	buttonSets := parseInputPart2(string(ReadInput(13, false)))
	allButtonHits := []int{}
	for _, set := range buttonSets {
		costA := set.Buttons[0].Cost
		costB := set.Buttons[1].Cost
		deltaAX := set.Buttons[0].DeltaX
		deltaAY := set.Buttons[0].DeltaY
		deltaBX := set.Buttons[1].DeltaX
		deltaBY := set.Buttons[1].DeltaY
		targetX := set.Target.X
		targetY := set.Target.Y

		numA := float64(targetX*deltaBY-targetY*deltaBX) / float64(deltaAX*deltaBY-deltaAY*deltaBX)
		numB := float64(deltaAX*targetY-deltaAY*targetX) / float64(deltaAX*deltaBY-deltaAY*deltaBX)

		if numA == float64(int(numA)) && numB == float64(int(numB)) &&
			numA >= 0{
			hitsCost := int(numA)*costA + int(numB)*costB
			allButtonHits = append(allButtonHits, hitsCost)
		}
	}

	totalCost := 0
	for _, cost := range allButtonHits {
		totalCost += cost
	}

	fmt.Println(totalCost)
}

func parseInput(input string) []ButtonSet {
	var buttonSets []ButtonSet
	sets := strings.Split(input, "\n\n")

	for _, set := range sets {
		var buttonSet ButtonSet
		lines := strings.Split(set, "\n")

		// target coords
		if matches := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindStringSubmatch(lines[2]); matches != nil {
			buttonSet.Target.X, _ = strconv.Atoi(matches[1])
			buttonSet.Target.Y, _ = strconv.Atoi(matches[2])
		}

		// button A coords
		if matches := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(lines[0]); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			buttonSet.Buttons = append(buttonSet.Buttons, Button{
				Point:  Point{X: 0, Y: 0},
				DeltaX: x,
				DeltaY: y,
				Cost:   3,
			})
		}

		// button B coords
		if matches := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(lines[1]); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			buttonSet.Buttons = append(buttonSet.Buttons, Button{
				Point:  Point{X: 0, Y: 0},
				DeltaX: x,
				DeltaY: y,
				Cost:   1,
			})
		}

		buttonSets = append(buttonSets, buttonSet)
	}

	return buttonSets
}

func parseInputPart2(input string) []ButtonSet {
	var buttonSets []ButtonSet
	sets := strings.Split(input, "\n\n")
	const offset = 10000000000000

	for _, set := range sets {
		var buttonSet ButtonSet
		lines := strings.Split(set, "\n")

		// target coords with offset
		if matches := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindStringSubmatch(lines[2]); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			buttonSet.Target.X = x + offset
			buttonSet.Target.Y = y + offset
		}

		// button A coords
		if matches := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(lines[0]); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			buttonSet.Buttons = append(buttonSet.Buttons, Button{
				Point:  Point{X: 0, Y: 0},
				DeltaX: x,
				DeltaY: y,
				Cost:   3,
			})
		}

		// button B coords
		if matches := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(lines[1]); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			buttonSet.Buttons = append(buttonSet.Buttons, Button{
				Point:  Point{X: 0, Y: 0},
				DeltaX: x,
				DeltaY: y,
				Cost:   1,
			})
		}

		buttonSets = append(buttonSets, buttonSet)
	}

	return buttonSets
}
