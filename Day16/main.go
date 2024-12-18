package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const SIZE = 141
const (
	UP int = iota
	LEFT
	DOWN
	RIGHT
)

func outGrid(grid *[SIZE][SIZE][2]int, pos [3]int, direction int) {

	for i := pos[0] - 1; i < pos[0]+2; i++ {
		for j := pos[1] - 1; j < pos[1]+2; j++ {
			if grid[i][j][0] == -1 {
				fmt.Printf("%8c", '#')
			} else {
				fmt.Printf("%8d", grid[i][j][1])
			}
		}
		fmt.Print("\n")

	}
	fmt.Print(direction, " ", pos[2], " ", "\n")
	time.Sleep(time.Millisecond * 1500)
	fmt.Print("\n\n")
}

func moveReindeer(grid *[SIZE][SIZE][2]int, pos [3]int, direction int) {

	dirCount := direction - pos[2]
	if dirCount < 0 {
		dirCount = 4 + dirCount
	}
	if dirCount > 2 {
		dirCount = 4 - dirCount
	}
	switch direction {
	case UP:
		if grid[pos[0]-1][pos[1]][0] == -1 {
			return
		} else {
			//outGrid(grid, pos, direction)
			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]-1][pos[1]][1] > newValue || grid[pos[0]-1][pos[1]][1] == -1 {
				grid[pos[0]-1][pos[1]][1] = newValue
				moveReindeer(grid, [3]int{pos[0] - 1, pos[1], UP}, RIGHT)
				moveReindeer(grid, [3]int{pos[0] - 1, pos[1], UP}, UP)
				moveReindeer(grid, [3]int{pos[0] - 1, pos[1], UP}, LEFT)
				return
			} else {
				return
			}
		}

	case DOWN:

		if grid[pos[0]+1][pos[1]][0] == -1 {
			return
		} else {
			//outGrid(grid, pos, direction)
			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]+1][pos[1]][1] > newValue || grid[pos[0]+1][pos[1]][1] == -1 {
				grid[pos[0]+1][pos[1]][1] = newValue
				moveReindeer(grid, [3]int{pos[0] + 1, pos[1], DOWN}, RIGHT)
				moveReindeer(grid, [3]int{pos[0] + 1, pos[1], DOWN}, DOWN)
				moveReindeer(grid, [3]int{pos[0] + 1, pos[1], DOWN}, LEFT)
				return
			} else {
				return
			}
		}

	case RIGHT:

		if grid[pos[0]][pos[1]+1][0] == -1 {
			return
		} else {
			//outGrid(grid, pos, direction)

			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]][pos[1]+1][1] > newValue || grid[pos[0]][pos[1]+1][1] == -1 {
				grid[pos[0]][pos[1]+1][1] = newValue
				moveReindeer(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, RIGHT)
				moveReindeer(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, UP)
				moveReindeer(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, DOWN)
				return
			} else {
				return
			}
		}

	case LEFT:

		if grid[pos[0]][pos[1]-1][0] == -1 {
			return
		} else {
			//outGrid(grid, pos, direction)

			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]][pos[1]-1][1] > newValue || grid[pos[0]][pos[1]-1][1] == -1 {
				grid[pos[0]][pos[1]-1][1] = newValue
				moveReindeer(grid, [3]int{pos[0], pos[1] - 1, LEFT}, UP)
				moveReindeer(grid, [3]int{pos[0], pos[1] - 1, LEFT}, DOWN)
				moveReindeer(grid, [3]int{pos[0], pos[1] - 1, LEFT}, LEFT)
				return
			} else {
				return
			}
		}
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var grid [SIZE][SIZE][2]int
	var startPos [2]int
	var endPos [2]int
	var lineNumber int
	fmt.Println(UP)

	for scanner.Scan() {
		line := scanner.Text()
		for j := range SIZE {
			grid[lineNumber][j][1] = -1
			if line[j] == '#' {
				grid[lineNumber][j][0] = -1
			} else if line[j] == 'S' {
				grid[lineNumber][j][1] = 0
				startPos[0] = lineNumber
				startPos[1] = j
			} else if line[j] == 'E' {
				endPos[0] = lineNumber
				endPos[1] = j
			}
		}
		lineNumber++
	}
	currentPos := [3]int{startPos[0], startPos[1], RIGHT}
	moveReindeer(&grid, currentPos, DOWN)
	moveReindeer(&grid, currentPos, RIGHT)
	moveReindeer(&grid, currentPos, LEFT)
	moveReindeer(&grid, currentPos, UP)

	/*
		for i := range SIZE {
			for j := range SIZE {
				fmt.Printf("%8d", grid[i][j][1])
			}
			fmt.Print("\n")
		}
	*/

	fmt.Println(grid[endPos[0]][endPos[1]][1])
}
