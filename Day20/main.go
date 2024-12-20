package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const SIZE = 141

//const SIZE = 15

func traverse(grid *[SIZE][SIZE]int, pos [2]int) {
	if pos[0]+1 < SIZE {
		if grid[pos[0]+1][pos[1]] != -1 {

			if grid[pos[0]+1][pos[1]] > grid[pos[0]][pos[1]]+1 || grid[pos[0]+1][pos[1]] == 0 {
				grid[pos[0]+1][pos[1]] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0] + 1, pos[1]})
			}
		}
	}
	if pos[0]-1 >= 0 {
		if grid[pos[0]-1][pos[1]] != -1 {

			if grid[pos[0]-1][pos[1]] > grid[pos[0]][pos[1]]+1 || grid[pos[0]-1][pos[1]] == 0 {
				grid[pos[0]-1][pos[1]] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0] - 1, pos[1]})
			}
		}
	}
	if pos[1]+1 < SIZE {
		if grid[pos[0]][pos[1]+1] != -1 {

			if grid[pos[0]][pos[1]+1] > grid[pos[0]][pos[1]]+1 || grid[pos[0]][pos[1]+1] == 0 {
				grid[pos[0]][pos[1]+1] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0], pos[1] + 1})
			}
		}
	}
	if pos[1]-1 >= 0 {
		if grid[pos[0]][pos[1]-1] != -1 {

			if grid[pos[0]][pos[1]-1] > grid[pos[0]][pos[1]]+1 || grid[pos[0]][pos[1]-1] == 0 {
				grid[pos[0]][pos[1]-1] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0], pos[1] - 1})
			}
		}
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var grid [SIZE][SIZE]int
	var startPos [2]int
	var endPos [2]int
	var pointsToRem [][2]int

	{
		lineNumber := 0
		for scanner.Scan() {
			line := scanner.Text()
			for j := range line {
				if line[j] == '#' {
					grid[lineNumber][j] = -1
					pointsToRem = append(pointsToRem, [2]int{lineNumber, j})
				} else if line[j] == 'S' {
					startPos[0] = lineNumber
					startPos[1] = j
				} else if line[j] == 'E' {
					endPos[0] = lineNumber
					endPos[1] = j
				}
			}
			lineNumber++
		}
	}

	empty := grid
	t := time.Now()
	traverse(&grid, startPos)
	fmt.Println(time.Since(t))
	cheatSave := 100
	count := 0

	t = time.Now()
	for _, val := range pointsToRem {
		tmp := empty
		tmp[val[0]][val[1]] = 0
		traverse(&tmp, startPos)

		if grid[endPos[0]][endPos[1]]-tmp[endPos[0]][endPos[1]] >= cheatSave {
			count++
		}
	}
	fmt.Println(time.Since(t))
	fmt.Println(count)
}
