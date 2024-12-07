package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

var DIR_UP = [2]int{0, -1}
var DIR_RIGHT = [2]int{1, 0}
var DIR_DOWN = [2]int{0, 1}
var DIR_LEFT = [2]int{-1, 0}

func part1() (int, int) {
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

	dimensions := [2]int{width, height}
	visited := make([]bool, width*height)
	visited[guardPos[0]+guardPos[1]*width] = true
	tracePath(guardPos, guardDir, dimensions, obstacles, visited)

	traceCount := 0
	for _, v := range visited {
		if v {
			traceCount++
		}
	}

	loopCount := 0
	for i, v := range visited {
		if v {
			x := i % width
			y := i / width
			newObstaclePos := [2]int{x, y}
			obstacles[newObstaclePos] = true
			// fmt.Printf("Tracing w/ %v\n", newObstaclePos)
			// fmt.Printf("\t%v\n", obstacles)
			newVisited := make([]bool, len(visited))
			// copy(newVisited, visited)
			isLoop := tracePath(guardPos, guardDir, dimensions, obstacles, newVisited)
			if isLoop {
				// fmt.Printf("FOUND LOOP with %v\n", newObstaclePos)
				loopCount++
			}
			delete(obstacles, newObstaclePos)
		}
	}
	return traceCount, loopCount
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

func updateVisited(visited []bool, guardPos [2]int, guardDir [2]int, obstaclePos [2]int, width int) (change bool) {
	for {
		if visited[guardPos[0]+(guardPos[1]*width)] == false {
			change = true
		}

		visited[guardPos[0]+(guardPos[1]*width)] = true

		guardPos[0] += guardDir[0]
		guardPos[1] += guardDir[1]

		if guardPos[0] == obstaclePos[0] && guardPos[1] == obstaclePos[1] {
			break
		}
	}
	return change
}

func tracePath(guardPos, guardDir, dimensions [2]int, obstacles map[[2]int]bool, visited []bool) (isLoop bool) {
	width := dimensions[0]
	height := dimensions[1]
	thisPath := make([][2]int, dimensions[0]*dimensions[1])
	pathLength := make([]int, 5)
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

		updateVisited(visited, guardPos, guardDir, obstacleCoord, width)
		if atEdge {
			isLoop = false
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

		// fmt.Printf("GuardPos: %v GuardDir: %v\n", guardPos, guardDir)
		// fmt.Printf("\tStartPos: %v StartDir: %v\n", startPos, startDir)
		thisCount := 0
		for _, v := range thisPath {
			if v != [2]int{0, 0} {
				thisCount++
			}
		}
		// lastVal := pathLength[len(pathLength)-1]
		// pathLength[0] = thisCount
		pathLength = append(pathLength, thisCount)
		pathLength = pathLength[1:]
		sameLength := true
		lastLength := pathLength[0]
		for _, v := range pathLength {
			if v != lastLength || lastLength == 0 {
				sameLength = false
				break
			}
			lastLength = v
		}
		// fmt.Println(sameLength)
		// fmt.Println(pathLength)
		if sameLength {
			isLoop = true
			break
		}

		// if thisPath[guardPos[0]+guardPos[1]*width] == guardDir { // if guardPos == startPos && guardDir == startDir {
		// 	isLoop = true
		// 	break
		// }
		thisPath[guardPos[0]+guardPos[1]*width] = guardDir
	}

	return isLoop
}
