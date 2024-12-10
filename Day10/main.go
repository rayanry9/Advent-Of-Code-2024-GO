package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 45

func getRoute(arr *[SIZE][SIZE]int, x int, y int, val int) int {
	total := 0
	if val == 10 {
		(*arr)[x][y] = 11
		return 1
	}
	if x-1 >= 0 {
		if (*arr)[x-1][y] == val {
			total += getRoute(arr, x-1, y, val+1)
		}
	}
	if x+1 < SIZE {
		if (*arr)[x+1][y] == val {
			total += getRoute(arr, x+1, y, val+1)
		}
	}
	if y-1 >= 0 {
		if (*arr)[x][y-1] == val {
			total += getRoute(arr, x, y-1, val+1)
		}
	}
	if y+1 < SIZE {
		if (*arr)[x][y+1] == val {
			total += getRoute(arr, x, y+1, val+1)
		}
	}
	return total
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var arr [SIZE][SIZE]int
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			arr[lineNumber][i] = int(line[i] - '0')
		}
		lineNumber++
	}

	total := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if arr[i][j] == 0 {
				temp := arr
				total += getRoute(&temp, i, j, 1)
			}
		}
	}

	fmt.Println(total)
}
