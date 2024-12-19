package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SIZE = 71

func traverse(grid *[SIZE][SIZE]int, pos [2]int) {
	if pos[0]+1 < SIZE {

		if grid[pos[0]+1][pos[1]] != -1 {
			if grid[pos[0]+1][pos[1]] > grid[pos[0]][pos[1]]+1 || grid[pos[0]+1][pos[1]] == -2 {
				grid[pos[0]+1][pos[1]] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0] + 1, pos[1]})
			}
		}

	}

	if pos[0]-1 >= 0 {
		if grid[pos[0]-1][pos[1]] != -1 {

			if grid[pos[0]-1][pos[1]] > grid[pos[0]][pos[1]]+1 || grid[pos[0]-1][pos[1]] == -2 {
				grid[pos[0]-1][pos[1]] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0] - 1, pos[1]})
			}
		}
	}

	if pos[1]+1 < SIZE {
		if grid[pos[0]][pos[1]+1] != -1 {
			if grid[pos[0]][pos[1]+1] > grid[pos[0]][pos[1]]+1 || grid[pos[0]][pos[1]+1] == -2 {
				grid[pos[0]][pos[1]+1] = grid[pos[0]][pos[1]] + 1
				traverse(grid, [2]int{pos[0], pos[1] + 1})
			}
		}
	}

	if pos[1]-1 >= 0 {
		if grid[pos[0]][pos[1]-1] != -1 {
			if grid[pos[0]][pos[1]-1] > grid[pos[0]][pos[1]]+1 || grid[pos[0]][pos[1]-1] == -2 {
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

	for i := range SIZE {
		for j := range SIZE {
			grid[i][j] = -2
		}
	}
	grid[0][0] = 0

	var lineNumber int
	for scanner.Scan() {
		line := scanner.Text()
		commaSplit := strings.Split(line, ",")
		x, _ := strconv.Atoi(commaSplit[0])
		y, _ := strconv.Atoi(commaSplit[1])
		grid[y][x] = -1
		tmp := grid
		traverse(&tmp, [2]int{0, 0})

		if tmp[SIZE-1][SIZE-1] == -2 {
			fmt.Println(x, y)
			break
		}
		//fmt.Println(lineNumber)

		lineNumber++
	}

}
