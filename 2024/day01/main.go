package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt")
	}
	defer file.Close()

	if len(os.Args) == 1 {
		fmt.Println("Specify part1 or part2")
	} else if os.Args[1] == "part1" {
		calcDistanceSum(file)
		fmt.Println(time.Since(start))
	} else if os.Args[1] == "part2" {
		calcSimilarityScore(file)
		fmt.Println(time.Since(start))
	}
}

func calcSimilarityScore(file *os.File) {
	scanner := bufio.NewScanner(file)
	left := make([]int, 1000)
	right := make(map[int]int, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		lVal, _ := strconv.Atoi(fields[0])
		left = append(left, lVal)
		rVal, _ := strconv.Atoi(fields[1])
		right[rVal] += 1
	}

	similarityScore := 0
	for _, lVal := range left {
		similarityScore += (lVal * right[lVal])
	}
	fmt.Println("Similarity score: ", similarityScore)
}

func calcDistanceSum(file *os.File) {
	scanner := bufio.NewScanner(file)
	left := make([]int, 1000)
	right := make([]int, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		lVal, _ := strconv.Atoi(fields[0])
		left = append(left, lVal)
		rVal, _ := strconv.Atoi(fields[1])
		right = append(right, rVal)
	}

	slices.Sort(left)
	slices.Sort(right)
	distanceSum := 0
	for i, lVal := range left {
		rVal := right[i]
		distance := rVal - lVal
		if distance < 0 {
			distance = -distance
		}
		distanceSum += distance
	}

	fmt.Println("Distance sum: ", distanceSum)
}
