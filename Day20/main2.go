package main

import (
	"bufio"
	"fmt"
	"math"
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

func findSmaller(grid *[SIZE][SIZE]int, pos [2]int, threshold int) int {
	radius := 20
	total := 0
	for i := pos[0] - radius; i < SIZE; i++ {
		if i < 0 {
			continue
		}
		if i > pos[0]+radius {
			break
		}
		for j := pos[1] - radius; j < SIZE; j++ {
			if j < 0 {
				continue
			}
			if j > pos[1]+radius {
				break
			}

			if int(math.Abs(float64(pos[1]-j)))+int(math.Abs(float64(pos[0]-i))) <= radius {

				if grid[i][j]-grid[pos[0]][pos[1]]-int(math.Abs(float64(pos[1]-j)))-int(math.Abs(float64(pos[0]-i))) >= threshold {
					//					fmt.Println(i, j)
					//					fmt.Println(grid[i][j], grid[pos[0]][pos[1]])
					total++
				}
			}
		}
	}
	return total
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

	traverse(&grid, startPos)

	grid[startPos[0]][startPos[1]] = 0
	cheatSave := 100

	/*
		for i := range SIZE {
			fmt.Println(grid[i])
		}
	*/

	total := 0
	for i := range SIZE {
		for j := range SIZE {
			if grid[i][j] != -1 {
				total += findSmaller(&grid, [2]int{i, j}, cheatSave)
			}
		}
	}
	fmt.Println(total)

}
