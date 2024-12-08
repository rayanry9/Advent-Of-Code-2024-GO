package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	const SIZE = 130

	var arr [SIZE][SIZE]int
	lineNumber := 0
	var pos [2]int
	for scanner.Scan() {
		line := scanner.Text()
		for j := 0; j < SIZE; j++ {
			if line[j] == '.' {
				arr[lineNumber][j] = 0
			} else if line[j] == '#' {
				arr[lineNumber][j] = 1
			} else if line[j] == '^' {
				pos[0] = lineNumber
				pos[1] = j
				arr[lineNumber][j] = 2
			}
		}
		lineNumber++
	}
	direction := [2]int{0, -1}

	for {
		if pos[0]+direction[1] >= SIZE || pos[0]+direction[1] < 0 || pos[1]+direction[0] >= SIZE || pos[1]+direction[0] < 0 {
			break
		}
		if arr[pos[0]+direction[1]][pos[1]+direction[0]] == 0 || arr[pos[0]+direction[1]][pos[1]+direction[0]] == 2 {
			pos[0] += direction[1]
			pos[1] += direction[0]
			arr[pos[0]][pos[1]] = 2
		} else if arr[pos[0]+direction[1]][pos[1]+direction[0]] == 1 {

			if direction[0] == 0 && direction[1] == -1 {
				direction[0] = 1
				direction[1] = 0
			} else if direction[0] == 1 && direction[1] == 0 {
				direction[0] = 0
				direction[1] = 1
			} else if direction[0] == 0 && direction[1] == 1 {
				direction[0] = -1
				direction[1] = 0
			} else if direction[0] == -1 && direction[1] == 0 {
				direction[0] = 0
				direction[1] = -1
			}
		}
	}

	count := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if arr[i][j] == 2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
