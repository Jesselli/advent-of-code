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
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	ordering := map[int][]int{}

	wellOrderedUpdates := make([]string, 0)
	reorderedUpdates := make([][]int, 0)
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
			pgs := csvToInts(line)
			legalOrder, badIdx, swapIdx := pgOrderLegal(pgs, ordering)
			if legalOrder {
				wellOrderedUpdates = append(wellOrderedUpdates, line)
			} else {
				newPgs := make([]int, len(pgs))
				copy(newPgs, pgs)
				for !legalOrder {
					newPgs[badIdx], newPgs[swapIdx] = newPgs[swapIdx], newPgs[badIdx]
					legalOrder, badIdx, swapIdx = pgOrderLegal(newPgs, ordering)
				}
				reorderedUpdates = append(reorderedUpdates, newPgs)
				// fmt.Printf("Found legal order:\n")
				// fmt.Printf("\tOld: %s:", line)
				// fmt.Printf("\tNew: %v:", newPgs)
				// fmt.Printf("%s, badIdx: %d, swapIdx: %d\n", line, badIdx, swapIdx)
			}
		}
	}

	// TODO: Lots of redundant code here w/r/t line spliting and converting to ints
	wellOrderedSum := 0
	for _, update := range wellOrderedUpdates {
		pages := strings.Split(update, ",")
		middlePage := pages[len(pages)/2]
		middlePgNum, _ := strconv.Atoi(middlePage)
		wellOrderedSum += middlePgNum
	}

	reorderedSum := 0
	for _, update := range reorderedUpdates {
		middlePage := update[len(update)/2]
		reorderedSum += middlePage
	}

	return wellOrderedSum, reorderedSum
}

func csvToInts(pgUpdatesLine string) []int {
	var err error
	split := strings.Split(pgUpdatesLine, ",")
	pgs := make([]int, len(split))
	for i, s := range split {
		pgs[i], err = strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("Unable to parse pg: %w", err))
		}
	}
	return pgs
}

func pgOrderLegal(pgUpdates []int, pgOrder map[int][]int) (bool, int, int) {
	inOrder := true
	badIdx := -1
	swapIdx := -1
	for i, pg := range pgUpdates {
		for j := 0; j < i; j++ {
			prevPg := pgUpdates[j]
			precedingPgs := pgOrder[pg]
			if slices.Index(precedingPgs, prevPg) != -1 {
				inOrder = false
				badIdx = i
				swapIdx = j
				break
			}
		}
		if !inOrder {
			break
		}
	}
	return inOrder, badIdx, swapIdx
}
