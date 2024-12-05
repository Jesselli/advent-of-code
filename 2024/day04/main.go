package main

import (
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	// part1()
	answer := part2()
	fmt.Printf("Answer: %d\nDuration: %v\n", answer, time.Since(start))
}

var shapes [4][3][3]byte

func init() {
	shapes[0] = [3][3]byte{{'M', '.', 'S'}, {'.', 'A', '.'}, {'M', '.', 'S'}}
	shapes[1] = [3][3]byte{{'M', '.', 'M'}, {'.', 'A', '.'}, {'S', '.', 'S'}}
	shapes[2] = [3][3]byte{{'S', '.', 'M'}, {'.', 'A', '.'}, {'S', '.', 'M'}}
	shapes[3] = [3][3]byte{{'S', '.', 'S'}, {'.', 'A', '.'}, {'M', '.', 'M'}}
}

func checkShapes(input []byte, lineLen int) int {
	count := 0
	for i := range input {
		if i%lineLen > lineLen-3 {
			continue
		}

		if (i + 2 + (lineLen * 2)) > len(input) {
			continue
		}

		for _, shape := range shapes {
			match := shape[0][0] == input[i]
			match = match && (shape[0][2] == input[i+2])
			match = match && (shape[1][1] == input[i+lineLen+1])
			match = match && (shape[2][0] == input[i+(lineLen*2)])
			match = match && (shape[2][2] == input[i+2+(lineLen*2)])

			if match {
				count++
			}
		}
	}

	return count
}

func part2() int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Couldn't open input")
	}

	lineLen := 0
	inputWithoutNewline := make([]byte, 0)
	for i, v := range input {
		if v == '\n' {
			if lineLen == 0 {
				lineLen = i
			}
			continue
		}
		inputWithoutNewline = append(inputWithoutNewline, v)
	}
	return checkShapes(inputWithoutNewline, lineLen)
}

func part1() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Couldn't open input")
	}

	lineLen := 0
	inputWithoutNewline := make([]byte, 0)
	for i, v := range input {
		if v == '\n' {
			if lineLen == 0 {
				lineLen = i
			}
			continue
		}
		inputWithoutNewline = append(inputWithoutNewline, v)
	}

	pattern := "XMAS"
	patternBytes := []byte(pattern)
	slices.Reverse(patternBytes)
	reverse := string(patternBytes)
	count := 0
	count += getMatchCount(inputWithoutNewline, lineLen, pattern, reverse, 1)         // Horizontal
	count += getMatchCount(inputWithoutNewline, lineLen, pattern, reverse, lineLen)   // Vertical
	count += getMatchCount(inputWithoutNewline, lineLen, pattern, reverse, lineLen-1) // Diagonal left
	count += getMatchCount(inputWithoutNewline, lineLen, pattern, reverse, lineLen+1) // Diagonal right
	fmt.Printf("Count: %d\n", count)
}

func getMatchCount(input []byte, lineLen int, pattern string, reverse string, delta int) int {
	count := 0
	inputLen := len(input)
	for i := range input {
		// Check bounds
		if i+delta >= inputLen || i+(2*delta) >= inputLen || i+(3*delta) >= inputLen {
			continue
		}

		wrapping := false
		if delta == lineLen-1 {
			wrapping = i%lineLen < len(pattern)-1
		} else if delta == lineLen+1 || delta == 1 {
			wrapping = i%lineLen > lineLen-len(pattern)
		} else if delta == lineLen {
			wrapping = false
		}

		match := input[i] == pattern[0]
		match = match && (input[i+delta] == pattern[1])
		match = match && (input[i+delta*2] == pattern[2])
		match = match && (input[i+delta*3] == pattern[3])
		if match && !wrapping {
			count++
		}

		rmatch := input[i] == reverse[0]
		rmatch = rmatch && (input[i+delta] == reverse[1])
		rmatch = rmatch && (input[i+delta*2] == reverse[2])
		rmatch = rmatch && (input[i+delta*3] == reverse[3])
		if rmatch && !wrapping {
			count++
		}
	}
	return count
}
