package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type equation struct {
	target   int
	testVals []int
}

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open input.txt")
	}
	scanner := bufio.NewScanner(file)
	equations := make([]equation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			msg := fmt.Errorf("Unable to parse target number: %w", err)
			panic(msg)
		}
		rightSplit := strings.Split(parts[1], " ")
		testValues := make([]int, len(rightSplit))
		for i, numStr := range rightSplit {
			testValues[i], err = strconv.Atoi(numStr)
			if err != nil {
				msg := fmt.Errorf("Unable to parse test value: %w", err)
				panic(msg)
			}
		}
		equations = append(equations, equation{target, testValues})
	}
	count := 0
	for _, eq := range equations {
		if len(eq.testVals) == 1 && eq.testVals[0] == eq.target {
			count++
			continue
		}

		if visit(eq, 1, eq.testVals[0]) {
			count += eq.target
		}
	}
	return count, 0
}

func visit(eq equation, idx int, currentVal int) bool {
	if idx == len(eq.testVals) && currentVal == eq.target {
		return true
	} else if idx >= len(eq.testVals) {
		return false
	}

	foundBySum := visit(eq, idx+1, currentVal+eq.testVals[idx])
	foundByMult := visit(eq, idx+1, currentVal*eq.testVals[idx])
	return foundBySum || foundByMult
}
