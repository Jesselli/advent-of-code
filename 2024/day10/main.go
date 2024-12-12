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

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Unable to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	trailMap := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		lineNums := make([]int, len(split))
		for x, v := range split {
			numVal, err := strconv.Atoi(v)
			if err != nil {
				panic("Could not parse numeric input")
			}
			lineNums[x] = numVal
		}
		trailMap = append(trailMap, lineNums)
	}
	// for _, v := range trailMap {
	// 	fmt.Println(v)
	// }

	// score := 0
	for y, row := range trailMap {
		for x, val := range row {
			if val != 0 {
				continue
			}
			fmt.Printf("Checking %d, %d\n", x, y)
			// found := make(map[vec2]bool)
			calcTrailHeadScore(trailMap, vec2{x, y}, vec2{0, 0})
		}
	}
	return count, 0
}

type vec2 [2]int

func (v vec2) add(vv vec2) vec2 {
	return vec2{v[0] + vv[0], v[1] + vv[1]}
}

var START = vec2{0, 0}
var RIGHT = vec2{1, 0}
var DOWN = vec2{0, 1}
var LEFT = vec2{-1, 0}
var UP = vec2{0, -1}

type trailMap [][]int

func (t trailMap) valAt(coord vec2) int {
	return t[coord[1]][coord[0]]
}

func (t trailMap) height() int {
	return len(t)
}

func (t trailMap) width() int {
	return len(t[0])
}

var count = 0

func calcTrailHeadScore(tMap trailMap, coord, lastDir vec2) bool {
	if tMap.valAt(coord) == 9 {
		count++
		return true
	}

	// score := 0
	width := tMap.width()
	height := tMap.height()
	thisVal := tMap.valAt(coord)
	if lastDir != RIGHT && coord[0] > 0 {
		leftCoord := coord.add(LEFT)
		leftVal := tMap.valAt(leftCoord)
		if leftVal == thisVal+1 {
			//fmt.Printf("Could move left from %v to %v!\n", coord, leftCoord)
			calcTrailHeadScore(tMap, leftCoord, LEFT)
			// if calcTrailHeadScore(tMap, leftCoord, LEFT, found) {
			// 	found[leftCoord] = true
			// }
		}
	}
	if lastDir != UP && coord[1] < height-1 {
		downCoord := coord.add(DOWN)
		downVal := tMap.valAt(downCoord)
		if downVal == thisVal+1 {
			//fmt.Printf("Could move down from %v to %v!\n", coord, downCoord)
			calcTrailHeadScore(tMap, downCoord, DOWN)
			// if calcTrailHeadScore(tMap, downCoord, DOWN, found) {
			// 	found[downCoord] = true
			// }
		}
	}
	if lastDir != DOWN && coord[1] > 0 {
		upCoord := coord.add(UP)
		upVal := tMap.valAt(upCoord)
		if upVal == thisVal+1 {
			//fmt.Printf("Could move up from %v to %v!\n", coord, upCoord)
			calcTrailHeadScore(tMap, upCoord, UP)
			// if calcTrailHeadScore(tMap, upCoord, UP, found) {
			// 	found[upCoord] = true
			// }
		}
	}
	if lastDir != LEFT && coord[0] < width-1 {
		rightCoord := coord.add(RIGHT)
		rightVal := tMap.valAt(rightCoord)
		if rightVal == thisVal+1 {
			//fmt.Printf("Could move right from %v to %v!\n", coord, rightCoord)
			calcTrailHeadScore(tMap, rightCoord, RIGHT)
			// if calcTrailHeadScore(tMap, rightCoord, RIGHT, found) {
			// 	found[rightCoord] = true
			// }
		}
	}

	return false
}
