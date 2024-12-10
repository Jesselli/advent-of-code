package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type equation struct {
	target   int
	testVals []int
}

type coord struct {
	x, y int
}

func (c1 coord) minus(c2 coord) coord {
	return coord{c1.x - c2.x, c1.y - c2.y}
}

func (c1 coord) plus(c2 coord) coord {
	return coord{c1.x + c2.x, c1.y + c2.y}
}

func (c coord) withinBounds(w, h int) bool {
	return c.x >= 0 && c.x < w && c.y >= 0 && c.y < h
}

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

	antennaLocs := make(map[rune][]coord)
	scanner := bufio.NewScanner(file)
	y := 0 // (0, 0) is top left
	width := -1
	for scanner.Scan() {
		line := scanner.Text()
		if width == -1 {
			width = len(line)
		}
		for x, ch := range line {
			if ch == '.' {
				continue
			}
			antennaLocs[ch] = append(antennaLocs[ch], coord{x, y})
		}
		y++
	}
	height := y

	antiNodes := make(map[coord]bool, 0)
	for _, v := range antennaLocs {
		for lIdx := 0; lIdx < len(v); lIdx++ {
			for rIdx := lIdx + 1; rIdx < len(v); rIdx++ {
				c1 := v[lIdx]
				c2 := v[rIdx]
				delta := c2.minus(c1)
				antiNodes[c1] = true
				antiNodes[c2] = true

				inBounds := true
				antiNode := c2
				for inBounds {
					antiNode = antiNode.plus(delta)
					inBounds = antiNode.withinBounds(width, height)
					if inBounds {
						antiNodes[antiNode] = true
					}
				}

				inBounds = true
				antiNode = c2
				for inBounds {
					antiNode = antiNode.minus(delta)
					inBounds = antiNode.withinBounds(width, height)
					if inBounds {
						antiNodes[antiNode] = true
					}
				}
				// antiNode1 := c2.plus(delta)
				// if antiNode1.withinBounds(width, height) {
				// 	antiNodes[antiNode1] = true
				// }
				// antiNode2 := c1.minus(delta)
				// if antiNode2.withinBounds(width, height) {
				// 	antiNodes[antiNode2] = true
				// }
			}
		}
	}

	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		if antiNodes[coord{x, y}] == true {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return len(antiNodes), 0
}
