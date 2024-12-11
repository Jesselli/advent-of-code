package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

func part1() (int, int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Unable to open input.txt")
	}

	// '0' -> 48 in DEC
	files := make([]int, 0)
	fileNum := 0
	for i, v := range input {
		itemLength := int(v - 48)
		if v == '\n' {
			continue
		}

		if i%2 == 0 {
			// file
			for j := 0; j < itemLength; j++ {
				files = append(files, fileNum)
			}
			fileNum++
		} else {
			// space -- represented by -1
			for j := 0; j < itemLength; j++ {
				files = append(files, -1)
			}
		}
	}

	lIdx := 0              // Points to space
	rIdx := len(files) - 1 // Points to file
	for lIdx != rIdx {
		if files[lIdx] != -1 {
			lIdx++
			continue
		}

		if files[rIdx] == -1 {
			rIdx--
			continue
		}

		files[lIdx], files[rIdx] = files[rIdx], files[lIdx]
	}

	answer := 0
	for i, v := range files {
		if v == -1 {
			break
		}

		answer += i * v
	}

	return answer, 0
}
