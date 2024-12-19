package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

type eq struct {
	x1 int
	x2 int
	x  int
	y1 int
	y2 int
	y  int
}

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Failed to open input.txt")
	}

	equations := make([]eq, 0)
	var currEq eq
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		if lineNum == 0 {
			currEq = eq{}

			x1 := splitLine[2][2 : len(splitLine[2])-1]
			x1Num, err := strconv.Atoi(x1)
			if err != nil {
				panic("Failed to parse x1")
			}
			currEq.x1 = x1Num

			y1 := splitLine[3][2:]
			y1Num, err := strconv.Atoi(y1)
			if err != nil {
				panic("Failed to parse y1")
			}
			currEq.y1 = y1Num
			lineNum++
		} else if lineNum == 1 {
			x2 := splitLine[2][2 : len(splitLine[2])-1]
			x2Num, err := strconv.Atoi(x2)
			if err != nil {
				panic("Failed to parse x2")
			}
			currEq.x2 = x2Num

			y2 := splitLine[3][2:]
			y2Num, err := strconv.Atoi(y2)
			if err != nil {
				panic("Failed to parse y2")
			}
			currEq.y2 = y2Num
			lineNum++
		} else if lineNum == 2 {
			x := splitLine[1][2 : len(splitLine[1])-1]
			xNum, err := strconv.Atoi(x)
			if err != nil {
				panic(err.Error())
			}
			currEq.x = xNum + 10000000000000

			y := splitLine[2][2:]
			yNum, err := strconv.Atoi(y)
			if err != nil {
				panic(err.Error())
			}
			currEq.y = yNum + 10000000000000
			lineNum++
			equations = append(equations, currEq)
		} else if lineNum == 3 {
			lineNum = 0
		}
	}

	answer := 0
	for _, e := range equations {
		a, b, hasSolution := calcA(e)
		if hasSolution {
			answer += a*3 + b
		}
	}

	return answer, 0
}

// Given two equations of the form
// x1*a + x2*b = X
// y1*a + y2*b = Y
// Solve for a and b
func calcA(e eq) (int, int, bool) {
	numerator := e.y2*e.x - e.x2*e.y
	denominator := e.x1*e.y2 - e.x2*e.y1
	a := numerator / denominator
	b := (e.y - e.y1*a) / e.y2

	hasSolution := true
	if (a*e.x1 + b*e.x2) != e.x {
		hasSolution = false
	} else if (a*e.y1 + b*e.y2) != e.y {
		hasSolution = false
	}

	return a, b, hasSolution
}
