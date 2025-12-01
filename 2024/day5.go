package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5_1() {
	data := strings.Split(string(ReadInput(5, false)), "\n\n")
	rulesData, pagesData := data[0], data[1:]

	rules := make([][]int, 0)
	pages := make([][]int, 0)
	legitRows := make([][]int, 0)

	// prepare rules [[47 53] [97 13] [97 61] ...
	for _, r := range strings.Split(rulesData, "\n") {
		rule := make([]int, 0)
		for _, numStr := range strings.Split(r, "|") {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			rule = append(rule, num) // append to rule, not rules
		}
		rules = append(rules, rule)
	}

	// prepare pages [[75 47 61 53 29] [97 61 53 29 13] ...
	for _, p := range strings.Split(pagesData[0], "\n") {
		page := make([]int, 0)
		for _, numStr := range strings.Split(p, ",") {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			page = append(page, num)
		}
		pages = append(pages, page)
	}

	validatePagesRow := func(rows [][]int) {
		// for each row
		for _, r := range rows {
			legit := false
			ruleFound := false

			// for each page in row
			for idx, p := range r {

				for _, rule := range rules {

					if rule[0] == p {
						// p on left so right must be later in row

						for idx2, p2 := range r {
							if p2 == rule[1] {
								ruleFound = true

								if idx2 > idx {
									legit = true
								} else {
									legit = false
									break
								}
							}
						}
					} else if rule[1] == p {
						// p on right so left must be earlier in row

						for idx2, p2 := range r {
							if p2 == rule[0] {
								ruleFound = true

								if idx2 < idx {
									legit = true
								} else {
									legit = false
									break
								}
							}
						}

					}

					// exit from page loop if found rule violation
					if !legit && ruleFound {
						break
					}
				}

				// exit from page loop if found rule with violation
				if !legit && ruleFound {
					break
				}
			}

			if legit {
				legitRows = append(legitRows, r)
			}
		}

	}

	validatePagesRow(pages)
	sum := 0
	for _, r := range legitRows {
		middleIndex := len(r) / 2
		sum += r[middleIndex]
	}

	fmt.Printf("Day5_1: %d\n", sum)
}

func Day5_2() {
	data := strings.Split(string(ReadInput(5, false)), "\n\n")
	rulesData, pagesData := data[0], data[1:]

	rules := make([][]int, 0)
	pages := make([][]int, 0)
	invalidRowsSorted := make([][]int, 0)
	// prepare rules [[47 53] [97 13] [97 61] ...
	for _, r := range strings.Split(rulesData, "\n") {
		rule := make([]int, 0)
		for _, numStr := range strings.Split(r, "|") {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			rule = append(rule, num) // append to rule, not rules
		}
		rules = append(rules, rule)
	}

	// prepare pages [[75 47 61 53 29] [97 61 53 29 13] ...
	for _, p := range strings.Split(pagesData[0], "\n") {
		page := make([]int, 0)
		for _, numStr := range strings.Split(p, ",") {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			page = append(page, num)
		}
		pages = append(pages, page)
	}

	sortRow := func(row []int) []int {
		// brute force try until rules applied
		for {
			swapped := false

			// for each page in row
			for p := 0; p < len(row)-1; p++ {

				// for each number after page
				for nextP := p + 1; nextP < len(row); nextP++ {

					// check rule match and swap if legit
					for _, rule := range rules {
						if row[nextP] == rule[0] && row[p] == rule[1] {
							row[p], row[nextP] = row[nextP], row[p]
							swapped = true
						}
					}
				}
			}
			if !swapped {
				break
			}
		}
		return row
	}
	validatePagesRow := func(rows [][]int) {
		// for each row
		for _, r := range rows {
			legit := true
			ruleFound := false

			// for each page in row
			for idx, p := range r {
				for _, rule := range rules {
					if rule[0] == p {
						// p on left so right must be later in row
						for idx2, p2 := range r {
							if p2 == rule[1] {
								ruleFound = true
								if idx2 <= idx {
									legit = false
									break
								}
							}
						}
					} else if rule[1] == p {
						// p on right so left must be earlier in row
						for idx2, p2 := range r {
							if p2 == rule[0] {
								ruleFound = true
								if idx2 >= idx {
									legit = false
									break
								}
							}
						}
					}

					if !legit && ruleFound {
						break
					}
				}

				if !legit && ruleFound {
					break
				}
			}

			if !legit {
				sortedRow := make([]int, len(r))
				copy(sortedRow, r)
				sortedRow = sortRow(sortedRow)
				invalidRowsSorted = append(invalidRowsSorted, sortedRow)
			}
		}
	}

	validatePagesRow(pages)
	sum := 0
	for _, r := range invalidRowsSorted {
		middleIndex := len(r) / 2
		sum += r[middleIndex]
	}

	fmt.Printf("Day5_1: %d\n", sum)
}
