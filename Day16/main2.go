package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const SIZE = 141

//const SIZE = 15

// const SIZE = 17
const (
	UP int = iota
	LEFT
	DOWN
	RIGHT
)

func outGrid2(grid *[SIZE][SIZE][3]int, pos [3]int, direction int) {

	for i := pos[0] - 1; i < pos[0]+2; i++ {
		if i < 0 || i >= SIZE {
			continue
		}
		for j := pos[1] - 1; j < pos[1]+2; j++ {
			if j < 0 || j >= SIZE {
				continue
			}
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

func moveReindeer2(grid *[SIZE][SIZE][3]int, pos [3]int, direction int) {

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
			//outGrid2(grid, pos, direction)
			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]-1][pos[1]][1] > newValue || grid[pos[0]-1][pos[1]][1] == -1 {
				grid[pos[0]-1][pos[1]][1] = newValue
				moveReindeer2(grid, [3]int{pos[0] - 1, pos[1], UP}, RIGHT)
				moveReindeer2(grid, [3]int{pos[0] - 1, pos[1], UP}, UP)
				moveReindeer2(grid, [3]int{pos[0] - 1, pos[1], UP}, LEFT)
				return
			} else {
				return
			}
		}

	case DOWN:

		if grid[pos[0]+1][pos[1]][0] == -1 {
			return
		} else {
			//outGrid2(grid, pos, direction)
			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]+1][pos[1]][1] > newValue || grid[pos[0]+1][pos[1]][1] == -1 {
				grid[pos[0]+1][pos[1]][1] = newValue
				moveReindeer2(grid, [3]int{pos[0] + 1, pos[1], DOWN}, RIGHT)
				moveReindeer2(grid, [3]int{pos[0] + 1, pos[1], DOWN}, DOWN)
				moveReindeer2(grid, [3]int{pos[0] + 1, pos[1], DOWN}, LEFT)
				return
			} else {
				return
			}
		}

	case RIGHT:

		if grid[pos[0]][pos[1]+1][0] == -1 {
			return
		} else {
			//outGrid2(grid, pos, direction)

			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]][pos[1]+1][1] > newValue || grid[pos[0]][pos[1]+1][1] == -1 {
				grid[pos[0]][pos[1]+1][1] = newValue
				moveReindeer2(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, RIGHT)
				moveReindeer2(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, UP)
				moveReindeer2(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, DOWN)
				return
			} else {
				return
			}
		}

	case LEFT:

		if grid[pos[0]][pos[1]-1][0] == -1 {
			return
		} else {
			//outGrid2(grid, pos, direction)

			newValue := grid[pos[0]][pos[1]][1] + 1 + (1000 * dirCount)
			if grid[pos[0]][pos[1]-1][1] > newValue || grid[pos[0]][pos[1]-1][1] == -1 {
				grid[pos[0]][pos[1]-1][1] = newValue
				moveReindeer2(grid, [3]int{pos[0], pos[1] - 1, LEFT}, UP)
				moveReindeer2(grid, [3]int{pos[0], pos[1] - 1, LEFT}, DOWN)
				moveReindeer2(grid, [3]int{pos[0], pos[1] - 1, LEFT}, LEFT)
				return
			} else {
				return
			}
		}
	}
}

func countGrid(grid *[SIZE][SIZE][3]int, pos [3]int, count *int, direction int, currVal int) {

	dirCount := direction - pos[2]
	if dirCount < 0 {
		dirCount = 4 + dirCount
	}
	if dirCount > 2 {
		dirCount = 4 - dirCount
	}
	currVal -= (1 + (1000 * dirCount))

	//fmt.Println(*count)
	//fmt.Println(currVal)
	//outGrid2(grid, [3]int{pos[0], pos[1], 0}, direction)

	switch direction {
	case DOWN:
		if grid[pos[0]+1][pos[1]][1] <= currVal && grid[pos[0]+1][pos[1]][0] != -1 {
			countGrid(grid, [3]int{pos[0] + 1, pos[1], DOWN}, count, DOWN, currVal)
			countGrid(grid, [3]int{pos[0] + 1, pos[1], DOWN}, count, LEFT, currVal)
			countGrid(grid, [3]int{pos[0] + 1, pos[1], DOWN}, count, RIGHT, currVal)
		}
	case UP:
		if grid[pos[0]-1][pos[1]][1] <= currVal && grid[pos[0]-1][pos[1]][0] != -1 {
			countGrid(grid, [3]int{pos[0] - 1, pos[1], UP}, count, UP, currVal)
			countGrid(grid, [3]int{pos[0] - 1, pos[1], UP}, count, LEFT, currVal)
			countGrid(grid, [3]int{pos[0] - 1, pos[1], UP}, count, RIGHT, currVal)
		}
	case RIGHT:
		if grid[pos[0]][pos[1]+1][1] <= currVal && grid[pos[0]][pos[1]+1][0] != -1 {
			countGrid(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, count, DOWN, currVal)
			countGrid(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, count, UP, currVal)
			countGrid(grid, [3]int{pos[0], pos[1] + 1, RIGHT}, count, RIGHT, currVal)
		}
	case LEFT:
		if grid[pos[0]][pos[1]-1][1] <= currVal && grid[pos[0]][pos[1]-1][0] != -1 {
			countGrid(grid, [3]int{pos[0], pos[1] - 1, LEFT}, count, LEFT, currVal)
			countGrid(grid, [3]int{pos[0], pos[1] - 1, LEFT}, count, DOWN, currVal)
			countGrid(grid, [3]int{pos[0], pos[1] - 1, LEFT}, count, UP, currVal)
		}
	}

	if grid[pos[0]][pos[1]][2] == 1 {
		return
	}
	(*count)++
	grid[pos[0]][pos[1]][2] = 1
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var grid [SIZE][SIZE][3]int
	var startPos [2]int
	var endPos [2]int
	var lineNumber int

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
	moveReindeer2(&grid, currentPos, DOWN)
	moveReindeer2(&grid, currentPos, RIGHT)
	moveReindeer2(&grid, currentPos, LEFT)
	moveReindeer2(&grid, currentPos, UP)
	/*
		for i := range SIZE {
			for j := range SIZE {
				fmt.Printf("%7d", grid[i][j][1])
			}
			fmt.Print("\n")
		}
	*/
	var count int

	countGrid(&grid, [3]int{endPos[0], endPos[1], DOWN}, &count, DOWN, grid[endPos[0]][endPos[1]][1])

	fmt.Println(count)
}
