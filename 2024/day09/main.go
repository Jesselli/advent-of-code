package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	answer1, answer2 := part1()
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\nDuration: %v\n", answer1, answer2, time.Since(start))
}

func part1() (int, int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Unable to open input.txt")
	}
	input = input[:len(input)-1]

	diskSize := 0
	for _, v := range input {
		diskSize += int(v - 48)
	}
	disk := make([]int, diskSize)

	fileLengths := make([]int, len(input)/2+1)
	fileNum := 0
	idx := 0
	for i, v := range input {
		itemLength := int(v - 48)
		if i%2 == 0 {
			// file
			for j := 0; j < itemLength; j++ {
				disk[idx] = fileNum
				idx++
			}
			fileLengths[fileNum] = itemLength
			fileNum++
		} else {
			// space -- represented by -1
			for j := 0; j < itemLength; j++ {
				disk[idx] = -1
				idx++
			}
		}
	}

	rIdx := len(disk) - 1 // Points to file
	for rIdx > 0 {
		if disk[rIdx] == -1 {
			rIdx--
			continue
		}

		fileNum := disk[rIdx]
		fileLength := fileLengths[disk[rIdx]]

		lIdx := 0
		for lIdx < rIdx {
			if disk[lIdx] != -1 {
				lIdx++
				continue
			}

			spaceLength := 0
			for i := 0; i < len(disk); i++ {
				if lIdx+i > len(disk)-1 {
					break
				}

				if disk[lIdx+i] != -1 {
					break
				} else {
					spaceLength++

				}
			}

			if fileLength <= spaceLength {
				for i := range fileLength {
					disk[lIdx+i] = fileNum
					disk[rIdx-i] = -1
				}
				break
			}
			lIdx++
		}

		// Move left to the next file
		rIdx -= fileLength
	}

	// // DEFRAG PT 1
	// lIdx := 0              // Points to space
	// rIdx := len(files) - 1 // Points to file
	// for lIdx != rIdx {
	// 	if files[lIdx] != -1 {
	// 		lIdx++
	// 		continue
	// 	}
	//
	// 	if files[rIdx] == -1 {
	// 		rIdx--
	// 		continue
	// 	}
	//
	// 	files[lIdx], files[rIdx] = files[rIdx], files[lIdx]
	// }
	// fmt.Printf("%v\n", files)

	answer := 0
	for i, v := range disk {
		if v == -1 {
			continue
		}

		answer += i * v
	}

	return answer, 0
}
