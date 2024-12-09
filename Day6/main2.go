package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 130
const UP = 2
const DOWN = 3
const RIGHT = 4
const LEFT = 5

func getDirection(direction *[2]int) int {

	if direction[0] == 0 && direction[1] == -1 {
		return UP
	} else if direction[0] == 0 && direction[1] == 1 {
		return DOWN
	} else if direction[0] == 1 && direction[1] == 0 {
		return RIGHT
	} else {
		return LEFT
	}

}

func runSimulation(arr [SIZE][SIZE]int, pos [2]int, x int, y int) bool {
	arr[x][y] = -1
	direction := [2]int{0, -1}

	for {
		if pos[0]+direction[1] >= SIZE || pos[0]+direction[1] < 0 || pos[1]+direction[0] >= SIZE || pos[1]+direction[0] < 0 {
			break
		}

		if arr[pos[0]+direction[1]][pos[1]+direction[0]] >= 0 {

			pos[0] += direction[1]
			pos[1] += direction[0]

			arr[pos[0]][pos[1]] = getDirection(&direction)

		} else if arr[pos[0]+direction[1]][pos[1]+direction[0]] < 0 {
			if arr[pos[0]+direction[1]][pos[1]+direction[0]] == -getDirection(&direction) {
				return true
			}
			arr[pos[0]+direction[1]][pos[1]+direction[0]] = -getDirection(&direction)

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
	return false
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var arr [SIZE][SIZE]int
	var arr2 [SIZE][SIZE]int
	lineNumber := 0
	var pos [2]int
	for scanner.Scan() {
		line := scanner.Text()
		for j := 0; j < SIZE; j++ {
			if line[j] == '.' {
				arr[lineNumber][j] = 0
			} else if line[j] == '#' {
				arr[lineNumber][j] = -1
			} else if line[j] == '^' {
				pos[0] = lineNumber
				pos[1] = j
			}
		}
		lineNumber++
	}
	count := 0

	for x := 0; x < SIZE; x++ {
		for y := 0; y < SIZE; y++ {
			if x == pos[0] && y == pos[1] {
				continue
			}
			if runSimulation(arr, pos, x, y) {
				arr2[x][y] = 1
			}
		}
	}
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if arr2[i][j] == 1 && arr[i][j] != -1 {
				count++
			}
		}
	}

	fmt.Println(count)
}
