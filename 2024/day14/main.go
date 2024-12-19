package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector [2]int

func (v *vector) add(vv vector) {
	v[0] += vv[0]
	v[1] += vv[1]
}

type robot struct {
	pos *vector
	vel *vector
}

func (r *robot) pacMove(w, h int) {
	r.pos.add(*r.vel)

	if r.pos[0] < 0 {
		r.pos[0] = w + r.pos[0]
	} else {
		r.pos[0] = r.pos[0] % w
	}

	if r.pos[1] < 0 {
		r.pos[1] = h + r.pos[1]
	} else {
		r.pos[1] = r.pos[1] % h
	}
}

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	robots := make([]robot, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		pos := parseVecStr(splitLine[0][2:])
		vel := parseVecStr(splitLine[1][2:])
		robot := robot{&pos, &vel}
		robots = append(robots, robot)
	}

	width := 101
	height := 103
	moveCount := 10000
	ii := 0
	for i := range moveCount {
		for _, robot := range robots {
			robot.pacMove(width, height)
		}
		if i == 584 || ii == 103 {
			displayMap(width, height, robots)
			fmt.Println(i)
			fmt.Scanln()
			ii = 0
		}
		ii++
	}

	quadCounts := [4]int{0, 0, 0, 0}
	for _, robot := range robots {
		if robot.pos[0] == width/2 || robot.pos[1] == height/2 {
			continue
		} else if robot.pos[0] < width/2 && robot.pos[1] < height/2 {
			quadCounts[0] += 1
		} else if robot.pos[0] > width/2 && robot.pos[1] < height/2 {
			quadCounts[1] += 1
		} else if robot.pos[0] < width/2 && robot.pos[1] > height/2 {
			quadCounts[2] += 1
		} else if robot.pos[0] > width/2 && robot.pos[1] > height/2 {
			quadCounts[3] += 1
		}
	}
	fmt.Println(quadCounts)

	return quadCounts[0] * quadCounts[1] * quadCounts[2] * quadCounts[3], 0
}

func parseVecStr(vecStr string) vector {
	vec := [2]int{-1, -1}
	commaIdx := strings.Index(vecStr, ",")
	x, err := strconv.Atoi(vecStr[:commaIdx])
	if err != nil {
		panic(err.Error())
	}
	vec[0] = x

	y, err := strconv.Atoi(vecStr[commaIdx+1:])
	if err != nil {
		panic(err.Error())
	}
	vec[1] = y

	return vec
}

func displayMap(w, h int, robots []robot) {
	fmt.Fprint(os.Stdout, "\x1b[1;1H")
	// fmt.Fprint(os.Stdout, "\x1b[31m")
	fmt.Print("\x1b[0J")
	// fmt.Print("\x1b[1;1H")
	for y := range h {
		for x := range w {
			hasRobo := false
			for _, r := range robots {
				if *r.pos == [2]int{x, y} {
					fmt.Printf("%c", 'x')
					hasRobo = true
					break
				}
			}
			if !hasRobo {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
}
