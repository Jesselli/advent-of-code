package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println("Couldn't open input")
	// }
	// defer file.Close()
	part2()
}

func part1() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Error reading input.txt")
	}
	mulRegExp := regexp.MustCompile(`mul\([0-9]*\,[0-9]*\)`)
	matches := mulRegExp.FindAllString(string(inputBytes), -1)
	fmt.Printf("Matches %v\n", matches)
	sum := 0
	for _, match := range matches {
		l, r := parseMatch(match)
		sum = sum + (l * r)
	}
	fmt.Println("Sum: ", sum)
}

func parseMatch(mulExp string) (int, int) {
	// mul(###,###)
	split := strings.Split(mulExp, ",")
	left := split[0][4:]
	right := split[1][0 : len(split[1])-1]
	fmt.Println(left, right)
	lNum, err := strconv.Atoi(left)
	if err != nil {
		panic("Could not parse left number")
	}
	rNum, err := strconv.Atoi(right)
	if err != nil {
		panic("Could not parse right number")
	}
	return lNum, rNum
}

func part2() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Error reading input.txt")
	}
	input := string(inputBytes)
	mulRegExp := regexp.MustCompile(`mul\([0-9]*\,[0-9]*\)`)
	mulMatches := mulRegExp.FindAllStringIndex(string(inputBytes), -1)
	doRegExp := regexp.MustCompile(`do\(\)`)
	doMatches := doRegExp.FindAllStringIndex(string(inputBytes), -1)
	dontRegExp := regexp.MustCompile(`don\'t\(\)`)
	dontMatches := dontRegExp.FindAllStringIndex(string(inputBytes), -1)
	sum := 0
	for _, v := range mulMatches {
		mulIdx := v[0]
		lastDontIdx := -1
		for _, v := range dontMatches {
			if v[0] < mulIdx {
				lastDontIdx = v[0]
			}
		}
		lastDoIdx := 0
		for _, v := range doMatches {
			if v[0] < mulIdx {
				lastDoIdx = v[0]
			}
		}
		if lastDontIdx > lastDoIdx {
			continue
		} else {
			x, y := parseMatch(input[v[0]:v[1]])
			sum = sum + (x * y)
		}
	}
	fmt.Println("Sum: ", sum)
}
