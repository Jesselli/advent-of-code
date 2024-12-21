package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type vec2 [2]int

type room struct {
	squares map[vec2]rune
	width   int
	height  int
	roboPos vec2
}

func initRoom() room {
	rm := room{}
	rm.squares = make(map[vec2]rune)
	rm.roboPos = vec2{-1, -1}
	return rm
}

func (r *room) moveRobo(dir vec2) {
	trgSq := r.roboPos.add(dir)
	canMove, boxesToMove := r.freeToMove(r.roboPos, dir)
	if canMove {
		r.roboPos = trgSq
		for _, b := range boxesToMove {
			r.squares[b] = '.'
		}
		for _, b := range boxesToMove {
			r.squares[b.add(dir)] = 'O'
		}
	}
}

func (r *room) freeToMove(sq vec2, dir vec2) (bool, []vec2) {
	canMove := true
	boxesToMove := make([]vec2, 0)

	nextSq := sq.add(dir)
	for nextSq[0] >= 0 && nextSq[0] < r.width && nextSq[1] >= 0 && nextSq[1] < r.height {
		if r.squares[nextSq] == 'O' {
			boxesToMove = append(boxesToMove, nextSq)
		} else if r.squares[nextSq] == '.' {
			canMove = true
			break
		} else if r.squares[nextSq] == '#' {
			canMove = false
			break
		}
		nextSq = nextSq.add(dir)
	}
	return canMove, boxesToMove
}

func (r *room) print() {
	for y := range r.height {
		for x := range r.width {
			if (r.roboPos == vec2{x, y}) {
				fmt.Printf("@")
			} else if sq, ok := r.squares[vec2{x, y}]; ok {
				fmt.Printf("%c", sq)
			}

		}
		fmt.Println()
	}
}

func (r *room) sumGPSCoords() int {
	sum := 0
	for k, v := range r.squares {
		if v == 'O' {
			sum += 100*k[1] + k[0]
		}
	}
	return sum
}

var (
	UP    = [2]int{0, -1}
	RIGHT = [2]int{1, 0}
	DOWN  = [2]int{0, 1}
	LEFT  = [2]int{-1, 0}
)

func (v vec2) add(vv vec2) vec2 {
	result := vec2{-1, -1}
	result[0] = v[0] + vv[0]
	result[1] = v[1] + vv[1]
	return result
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

	room := initRoom()
	roboInput := make([]vec2, 0)
	scanner := bufio.NewScanner(file)
	parsingRoom := true

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRoom = false
			continue
		}

		if parsingRoom {
			parseRoomInput(&room, y, line)
			// if x := strings.IndexRune(line, '@'); x != -1 {
			// 	room.roboPos[0] = x
			// 	room.roboPos[1] = y
			// }
			room.width = len(line)
			y++
		} else {
			input := parseRobotInput(line)
			roboInput = append(roboInput, input...)
		}
	}
	room.height = y
	// room.print()

	for _, d := range roboInput {
		// fmt.Scanln()
		// fmt.Println(d)
		room.moveRobo(d)
		// room.print()
	}
	room.print()

	return room.sumGPSCoords(), 0
}
func parseRoomInput(rm *room, y int, line string) {
	for x, r := range line {
		if r == '#' {
			// Wall
			rm.squares[vec2{x, y}] = r
		} else if r == 'O' {
			// Box
			rm.squares[vec2{x, y}] = r
		} else if r == '@' {
			// Robot
			rm.roboPos = vec2{x, y}
			rm.squares[vec2{x, y}] = '.'
		} else if r == '.' {
			rm.squares[vec2{x, y}] = r
		}
	}
}

func parseRobotInput(line string) []vec2 {
	result := make([]vec2, len(line))
	for i, r := range line {
		if r == '^' {
			result[i] = UP
		} else if r == '>' {
			result[i] = RIGHT
		} else if r == 'v' {
			result[i] = DOWN
		} else if r == '<' {
			result[i] = LEFT
		}
	}
	return result
}
