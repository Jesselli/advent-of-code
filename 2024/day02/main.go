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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Couldn't open input")
	}
	defer file.Close()
	part2(file)
}

func checkDelta(fields []string) bool {
	minDelta := 1
	maxDelta := 3
	numFields := len(fields)
	increasing := true
	for i := 0; i < numFields-1; i++ {
		this, _ := strconv.Atoi(fields[i])
		next, _ := strconv.Atoi(fields[i+1])
		delta := next - this
		if i == 0 {
			increasing = delta > 0
		}
		absDelta := delta
		if absDelta < 0 {
			absDelta = -absDelta
		}
		if absDelta < minDelta || absDelta > maxDelta {
			return false
		} else if increasing != (delta > 0) {
			return false
		}
	}
	return true
}

func part2(file *os.File) {
	start := time.Now()
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if checkDelta(fields) {
			safeCount++
		} else {
			for i := 0; i < len(fields); i++ {
				fieldsRevised := make([]string, 0)
				for j, v := range fields {
					if j != i {
						fieldsRevised = append(fieldsRevised, v)
					}
				}
				if checkDelta(fieldsRevised) {
					safeCount++
					break
				}
			}
		}
	}
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println("safeCount w/ fault tolerance: ", safeCount)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		minDelta := 1
		maxDelta := 3
		increasing := true
		numFields := len(fields)
		isSafe := true
		for i := 0; i < numFields-1; i++ {
			this, _ := strconv.Atoi(fields[i])
			next, _ := strconv.Atoi(fields[i+1])

			if i == 0 {
				increasing = next > this
			}

			delta := next - this
			if delta < 0 {
				if increasing {
					isSafe = false
					break
				}
				delta = -delta
			} else {
				if !increasing {
					isSafe = false
					break
				}
			}
			if delta < minDelta || delta > maxDelta {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}
	fmt.Println("safeCount: ", safeCount)
}
