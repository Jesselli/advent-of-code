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

	visited := make(map[coord]bool)
	cost := 0
	for y, row := range rows {
		for x, ch := range row {
			p, a := paintRegion([2]int{x, y}, ch, rows, &visited)
			cost += p * a
		}
	}
	fmt.Println("Cost:", cost)

	return 0, 0
}

type coord [2]int

var UP coord = [2]int{0, -1}
var RIGHT coord = [2]int{1, 0}
var DOWN coord = [2]int{0, 1}
var LEFT coord = [2]int{-1, 0}

func (c coord) add(cc coord) coord {
	return [2]int{c[0] + cc[0], c[1] + cc[1]}
}

func paintRegion(c coord, regionCh byte, rows [][]byte, visited *map[coord]bool) (perimeter, area int) {
	if c[0] < 0 || c[0] >= len(rows[0]) || c[1] < 0 || c[1] >= len(rows) {
		return 1, 0
	}

	ch := rows[c[1]][c[0]]
	if ch != regionCh {
		return 1, 0
	}

	if (*visited)[c] == true {
		return 0, 0
	}

	(*visited)[c] = true
	pU, aU := paintRegion(c.add(UP), regionCh, rows, visited)
	pR, aR := paintRegion(c.add(RIGHT), regionCh, rows, visited)
	pD, aD := paintRegion(c.add(DOWN), regionCh, rows, visited)
	pL, aL := paintRegion(c.add(LEFT), regionCh, rows, visited)

	return (pU + pR + pD + pL), (aU + aR + aD + aL) + 1
}
