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
	answer := part1()
	fmt.Printf("Answer: %d\nDuration: %v\n", answer, time.Since(start))
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	ordering := map[int][]int{}

	wellOrderedUpdates := make([]string, 0)
	parsingPageOrder := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingPageOrder = false
			continue
		}

		if parsingPageOrder {
			parts := strings.Split(line, "|")
			leftPg, err := strconv.Atoi(parts[0])
			if err != nil {
				errMsg := err.Error()
				panic(errMsg)
			}
			rightPg, err := strconv.Atoi(parts[1])
			if err != nil {
				panic("Couldn't parse right page num")
			}

			pgPrecedes := ordering[leftPg]
			ordering[leftPg] = append(pgPrecedes, rightPg)
		} else {
			// Parsing page updates
			badOrder := false
			pages := strings.Split(line, ",")
			for i, pg := range pages {
				currPg, err := strconv.Atoi(pg)
				if err != nil {
					panic(err.Error())
				}
				for j := 0; j < i; j++ {
					// Check past pages for ordering
					prevPg, _ := strconv.Atoi(pages[j])
					precedingPgs := ordering[currPg]
					k := slices.Index(precedingPgs, prevPg)
					if k != -1 {
						badOrder = true
						break
					}
				}
				if badOrder {
					break
				}
			}
			if !badOrder {
				fmt.Println("Good order: ", line)
				wellOrderedUpdates = append(wellOrderedUpdates, line)
			}
		}
	}

	// TODO: Lots of redundant code here w/r/t line spliting and converting to ints
	sum := 0
	for _, update := range wellOrderedUpdates {
		pages := strings.Split(update, ",")
		middlePage := pages[len(pages)/2]
		middlePgNum, _ := strconv.Atoi(middlePage)
		sum += middlePgNum
	}

	return sum
}
