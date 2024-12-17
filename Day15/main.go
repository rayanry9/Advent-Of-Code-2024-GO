package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 50

type MoveType int

const (
	UP MoveType = iota
	DOWN
	LEFT
	RIGHT
)

func moveRec(grid *[SIZE][SIZE]int, to [2]int, move MoveType) bool {

	switch move {
	case UP:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]+1][to[1]]
			(*grid)[to[0]+1][to[1]] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 {
			if moveRec(grid, [2]int{to[0] - 1, to[1]}, move) {
				(*grid)[to[0]][to[1]] = (*grid)[to[0]+1][to[1]]
				(*grid)[to[0]+1][to[1]] = 0
				return true
			} else {
				return false
			}
		}

	case DOWN:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]-1][to[1]]
			(*grid)[to[0]-1][to[1]] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 {
			if moveRec(grid, [2]int{to[0] + 1, to[1]}, move) {
				(*grid)[to[0]][to[1]] = (*grid)[to[0]-1][to[1]]
				(*grid)[to[0]-1][to[1]] = 0
				return true
			} else {
				return false
			}
		}

	case LEFT:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]+1]
			(*grid)[to[0]][to[1]+1] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 {
			if moveRec(grid, [2]int{to[0], to[1] - 1}, move) {
				(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]+1]
				(*grid)[to[0]][to[1]+1] = 0
				return true
			} else {
				return false
			}
		}

	case RIGHT:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]-1]
			(*grid)[to[0]][to[1]-1] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 {
			if moveRec(grid, [2]int{to[0], to[1] + 1}, move) {

				(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]-1]
				(*grid)[to[0]][to[1]-1] = 0
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var grid [SIZE][SIZE]int
	var pos [2]int

	for i := range SIZE {
		scanner.Scan()
		line := scanner.Text()
		for j := range SIZE {
			if line[j] == '#' {
				grid[i][j] = -1
			} else if line[j] == 'O' {
				grid[i][j] = 1
			} else if line[j] == '@' {
				grid[i][j] = 3
				pos[0] = i
				pos[1] = j
			}
		}
	}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		for j := range len(line) {
			if line[j] == '<' {
				if moveRec(&grid, [2]int{pos[0], pos[1] - 1}, LEFT) {
					pos[1]--
				}
			} else if line[j] == '>' {
				if moveRec(&grid, [2]int{pos[0], pos[1] + 1}, RIGHT) {
					pos[1]++
				}
			} else if line[j] == '^' {
				if moveRec(&grid, [2]int{pos[0] - 1, pos[1]}, UP) {
					pos[0]--
				}
			} else if line[j] == 'v' {
				if moveRec(&grid, [2]int{pos[0] + 1, pos[1]}, DOWN) {
					pos[0]++
				}
			}
		}
	}

	var value int
	for i := range SIZE {
		for j := range SIZE {
			if grid[i][j] == 1 {
				value += (100 * i) + j
			}
		}
	}
	fmt.Println(value)
}
