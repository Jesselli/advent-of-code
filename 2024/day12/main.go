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

type border uint8

const (
	BORDER_NONE border = 0b0000
	BORDER_U    border = 0b0001
	BORDER_R    border = 0b0010
	BORDER_D    border = 0b0100
	BORDER_L    border = 0b1000
)

func part1() (int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	rows := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowStr := scanner.Text()
		rows = append(rows, []byte(rowStr))
	}

	globalVisited := make(map[coord]border)
	costPt1 := 0
	costPt2 := 0
	for y, row := range rows {
		for x, ch := range row {
			if _, ok := globalVisited[coord{x, y}]; ok {
				continue
			}

			localVisited := make(map[coord]border)
			p, a := paintRegion([2]int{x, y}, [2]int{0, 0}, ch, rows, &localVisited)
			for k, v := range localVisited {
				globalVisited[k] = v
			}
			sides := calcSideCount(localVisited)

			fmt.Printf("%v\n", localVisited)
			fmt.Println("Sidecount:", sides)
			fmt.Println()

			costPt1 += p * a
			costPt2 += sides * a
		}
	}
	// fmt.Println("Cost:", costPt1)

	return costPt1, costPt2
}

func calcSideCount(visited map[coord]border) int {
	sideCount := 0
	minX, minY := 1000, 1000
	maxX, maxY := 0, 0 // HACK: Hard-coded values
	for k := range visited {
		if k[0] < minX {
			minX = k[0]
		}
		if k[0] > maxX {
			maxX = k[0]
		}
		if k[1] < minY {
			minY = k[1]
		}
		if k[1] > maxY {
			maxY = k[1]
		}
	}

	for y := minY; y < maxY+1; y++ {
		prevBorder := BORDER_NONE
		for x := minX; x < maxX+1; x++ {
			if b, ok := visited[coord{x, y}]; ok {
				if prevBorder&BORDER_U != BORDER_U && b&BORDER_U == BORDER_U {
					sideCount++
				}
				if prevBorder&BORDER_D != BORDER_D && b&BORDER_D == BORDER_D {
					sideCount++
				}
				prevBorder = b
			} else {
				prevBorder = BORDER_NONE
			}
		}
	}

	for x := minX; x < maxX+1; x++ {
		prevBorder := BORDER_NONE
		for y := minY; y < maxY+1; y++ {
			if b, ok := visited[coord{x, y}]; ok {
				if prevBorder&BORDER_L != BORDER_L && b&BORDER_L == BORDER_L {
					sideCount++
				}
				if prevBorder&BORDER_R != BORDER_R && b&BORDER_R == BORDER_R {
					sideCount++
				}
				prevBorder = b
			} else {
				prevBorder = BORDER_NONE
			}
		}
	}

	return sideCount
}

type coord [2]int

var UP coord = [2]int{0, -1}
var RIGHT coord = [2]int{1, 0}
var DOWN coord = [2]int{0, 1}
var LEFT coord = [2]int{-1, 0}

func (c coord) add(cc coord) coord {
	return [2]int{c[0] + cc[0], c[1] + cc[1]}
}

func paintRegion(c coord, dir coord, regionCh byte, rows [][]byte, visited *map[coord]border) (perimeter, area int) {
	newCoord := c.add(dir)
	if newCoord[0] < 0 {
		(*visited)[c] |= BORDER_L
		return 1, 0
	} else if newCoord[0] >= len(rows[0]) {
		(*visited)[c] |= BORDER_R
		return 1, 0
	} else if newCoord[1] < 0 {
		(*visited)[c] |= BORDER_U
		return 1, 0
	} else if newCoord[1] >= len(rows) {
		(*visited)[c] |= BORDER_D
		return 1, 0
	}

	ch := rows[newCoord[1]][newCoord[0]]
	if ch != regionCh {
		if dir == UP {
			(*visited)[c] |= BORDER_U
		} else if dir == RIGHT {
			(*visited)[c] |= BORDER_R
		} else if dir == DOWN {
			(*visited)[c] |= BORDER_D
		} else if dir == LEFT {
			(*visited)[c] |= BORDER_L
		}
		return 1, 0
	}

	if _, ok := (*visited)[newCoord]; ok {
		return 0, 0
	}

	(*visited)[newCoord] = BORDER_NONE
	pU, aU := paintRegion(newCoord, UP, regionCh, rows, visited)
	pR, aR := paintRegion(newCoord, RIGHT, regionCh, rows, visited)
	pD, aD := paintRegion(newCoord, DOWN, regionCh, rows, visited)
	pL, aL := paintRegion(newCoord, LEFT, regionCh, rows, visited)

	return (pU + pR + pD + pL), (aU + aR + aD + aL) + 1
}
