package main

import (
	"fmt"
	"strconv"
)

func Day9() {
	disk := ReadInput(9, false)

	var fileID int
	var decoded []int
	isFile := true

	for _, num := range disk {
		count, _ := strconv.Atoi(string(num))
		if isFile {
			// be hyperfocus on example, realize id can be multidigit number because of the meme
			for i := 0; i < count; i++ {
				decoded = append(decoded, fileID)
			}
			fileID++
			isFile = false
		} else {
			for i := 0; i < count; i++ {
				decoded = append(decoded, -1)
			}
			isFile = true
		}
	}

	sum := 0
	l, r := 0, len(decoded)-1
	for {
		// find next dot from left
		for l < len(decoded) && decoded[l] >= 0 {
			sum += decoded[l] * l
			l++
		}

		// find next number from right
		for r > l && decoded[r] < 0 {
			r--
		}

		if l >= r {
			break
		}

		decoded[l] = decoded[r]
		decoded[r] = -1
	}

	fmt.Println(sum)
}
