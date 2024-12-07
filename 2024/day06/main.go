package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	answer := part1()
	fmt.Printf("Answer 1: %d\nDuration: %v\n", answer, time.Since(start))
}

var DIR_UP = [2]int{0, -1}
var DIR_RIGHT = [2]int{1, 0}
var DIR_DOWN = [2]int{0, 1}
var DIR_LEFT = [2]int{-1, 0}

func part1() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}
	defer inputFile.Close()

	obstacles := make(map[[2]int]bool, 0)
	guardPos := [2]int{-1, -1}
	guardDir := [2]int{0, -1}
	width := 0
	height := 0
	scanner := bufio.NewScanner(inputFile)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for i := range line {
			ch := line[i]
			coord := [2]int{i, lineNum}
			if ch == '#' {
				obstacles[coord] = true
			} else if ch == '^' {
				guardPos = coord
			}
		}
		lineNum++
	}
	height = lineNum

	visited := make([]bool, width*height)
	visited[guardPos[0]+guardPos[1]*width] = true
	for guardPos[0] < width && guardPos[1] < height {
		obstacleCoord := [2]int{-1, -1}
		obstacleCoord = findFirstObstacle(guardDir, guardPos, obstacles)

		atEdge := false
		if obstacleCoord[0] > width {
			obstacleCoord[0] = width
			obstacleCoord[1] = guardPos[1]
			atEdge = true
		} else if obstacleCoord[1] > height {
			obstacleCoord[0] = guardPos[0]
			obstacleCoord[1] = height
			atEdge = true
		} else if obstacleCoord[0] < 0 {
			obstacleCoord[0] = 0
			obstacleCoord[1] = guardPos[1]
			atEdge = true
		} else if obstacleCoord[1] < 0 {
			obstacleCoord[0] = guardPos[0]
			obstacleCoord[1] = 0
			atEdge = true
		}

		visited = updateVisited(visited, guardPos, guardDir, obstacleCoord, width)
		if atEdge {
			break
		}

		if guardDir == DIR_UP {
			guardPos[1] = obstacleCoord[1] + 1
			guardDir = DIR_RIGHT
		} else if guardDir == DIR_RIGHT {
			guardPos[0] = obstacleCoord[0] - 1
			guardDir = DIR_DOWN
		} else if guardDir == DIR_DOWN {
			guardPos[1] = obstacleCoord[1] - 1
			guardDir = DIR_LEFT
		} else if guardDir == DIR_LEFT {
			guardPos[0] = obstacleCoord[0] + 1
			guardDir = DIR_UP
		}

		tempCount := 0
		for _, v := range visited {
			if v {
				tempCount++
			}
		}
		fmt.Println(obstacleCoord, guardPos, tempCount)
	}

	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}
	return count
}

func findFirstObstacle(guardDir, guardPos [2]int, obstacles map[[2]int]bool) [2]int {
	// FIX: Hard-Coded
	firstObstacle := [2]int{guardDir[0] * 1000, guardDir[1] * 1000}
	for coord := range obstacles {
		// (0,0) is top left
		if guardDir == DIR_UP {
			if coord[0] == guardPos[0] && coord[1] < guardPos[1] && coord[1] > firstObstacle[1] {
				firstObstacle = coord
			}
		} else if guardDir == DIR_RIGHT {
			if coord[0] > guardPos[0] && coord[1] == guardPos[1] && coord[0] < firstObstacle[0] {
				firstObstacle = coord
			}
		} else if guardDir == DIR_DOWN {
			if coord[0] == guardPos[0] && coord[1] > guardPos[1] && coord[1] < firstObstacle[1] {
				firstObstacle = coord
			}
		} else if guardDir == DIR_LEFT {
			if coord[0] < guardPos[0] && coord[1] == guardPos[1] && coord[0] > firstObstacle[0] {
				firstObstacle = coord
			}
		}
	}
	return firstObstacle
}

func updateVisited(visited []bool, guardPos [2]int, guardDir [2]int, obstaclePos [2]int, width int) []bool {
	for {
		visited[guardPos[0]+(guardPos[1]*width)] = true

		guardPos[0] += guardDir[0]
		guardPos[1] += guardDir[1]

		if guardPos[0] == obstaclePos[0] && guardPos[1] == obstaclePos[1] {
			break
		}
	}
	return visited
}
