package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func init() {
	newCache = make(map[[2]int]int)
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
	scanner.Split(bufio.ScanWords)
	rockNums := make([]int, 0)
	for scanner.Scan() {
		num := scanner.Text()
		val, err := strconv.Atoi(num)
		if err != nil {
			panic("Failed to parse rock number")
		}
		rockNums = append(rockNums, val)
	}

	count := 0
	for _, num := range rockNums {
		count += cacheBlink(num, 0, 75)
	}

	return count, 0
}

var newCache map[[2]int]int

func cacheBlink(rockNum, currentDepth, maxDepth int) (count int) {
	currTuple := [2]int{rockNum, currentDepth}
	if val, ok := newCache[currTuple]; ok {
		return val
	}

	left := -1
	right := -1
	count = 1

	if rockNum == 0 {
		// RULE 1: 0 -> 1
		left = 1
	} else if numDigits(rockNum)%2 == 0 {
		// RULE 2: Split nums with even # of digits
		left, right = splitNum(rockNum)
	} else {
		// RULE 3: If not Rule 1 or Rule 2, multiply by 2024
		left = rockNum * 2024
	}

	if currentDepth < maxDepth {
		// for range currentDepth {
		// 	fmt.Print("\t")
		// }
		// fmt.Printf("%2d Adding left %d\n", currentDepth, left)
		count = cacheBlink(left, currentDepth+1, maxDepth)
		if right != -1 {
			// for range currentDepth {
			// 	fmt.Print("\t")
			// }
			// fmt.Printf("%2d Adding right %d\n", currentDepth, right)
			count += cacheBlink(right, currentDepth+1, maxDepth)
		}
		newCache[currTuple] = count
	}

	return count
}

func numDigits(num int) int {
	digitCount := 0
	remaining := num
	for remaining != 0 {
		digitCount++
		remaining = remaining / 10
	}
	return digitCount
}

func splitNum(num int) (int, int) {
	numStr := strconv.Itoa(num)
	left := numStr[:len(numStr)/2]
	leftNum, err := strconv.Atoi(left)
	if err != nil {
		panic("Failed to parse left number")
	}
	right := numStr[len(numStr)/2:]
	rightNum, err := strconv.Atoi(right)
	if err != nil {
		panic("Failed to parse right number")
	}
	return leftNum, rightNum
}
