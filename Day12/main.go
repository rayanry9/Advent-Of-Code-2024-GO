package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const SIZE = 140

//const SIZE = 10

func getVal(i int, j int) int {
	return (i * SIZE) + j
}
func getCoords(val int) (int, int) {
	return val / SIZE, val % SIZE
}

func findSet(arr *[SIZE * SIZE]int, val int) int {
	if val == (*arr)[val] {
		return val
	} else {
		(*arr)[val] = findSet(arr, (*arr)[val])
		return (*arr)[val]
	}
}
func unionSets(arr *[SIZE * SIZE]int, a int, b int) {
	a = findSet(arr, a)
	b = findSet(arr, b)
	if a != b {
		if a < b {
			(*arr)[b] = a
		} else {
			(*arr)[a] = b
		}
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var arr [SIZE][SIZE]int
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		for idx, val := range line {
			arr[lineNumber][idx] = int(val)
		}
		lineNumber++
	}
	var parent [SIZE * SIZE]int
	var edges [SIZE][SIZE]int

	for i := range SIZE {
		for j := range SIZE {
			parent[getVal(i, j)] = getVal(i, j)
			edges[i][j] = 4
			if j-1 >= 0 {
				if arr[i][j] == arr[i][j-1] {
					edges[i][j]--
				}
			}
			if j+1 < SIZE {
				if arr[i][j] == arr[i][j+1] {
					edges[i][j]--
				}
			}
			if i-1 >= 0 {
				if arr[i][j] == arr[i-1][j] {
					edges[i][j]--
				}
			}
			if i+1 < SIZE {
				if arr[i][j] == arr[i+1][j] {
					edges[i][j]--
				}
			}
		}
	}
	for i := range SIZE {
		for j := range SIZE {
			if j == 0 {
				continue
			}
			if arr[i][j] == arr[i][j-1] {
				unionSets(&parent, getVal(i, j), getVal(i, j-1))
			}
		}
	}
	for i := range SIZE {
		if i == 0 {
			continue
		}
		for j := range SIZE {
			if arr[i][j] == arr[i-1][j] {
				unionSets(&parent, getVal(i, j), getVal(i-1, j))
			}
		}
	}

	var value int

	var perim [SIZE * SIZE]int
	var area [SIZE * SIZE]int
	counters := make([]int, 0)

	for i := range SIZE {
		for j := range SIZE {
			perim[findSet(&parent, getVal(i, j))] += edges[i][j]
			area[findSet(&parent, getVal(i, j))]++

			if !slices.Contains(counters, findSet(&parent, getVal(i, j))) {
				counters = append(counters, findSet(&parent, getVal(i, j)))
			}
		}
	}
	for _, val := range counters {
		value += perim[val] * area[val]
	}
	/*

		for i := range SIZE {
			for j := range SIZE {
				fmt.Printf("%6d", findSet(&parent, getVal(i, j)))
			}
			fmt.Print("\n")
		}
		for i := range SIZE {
			for j := range SIZE {
				fmt.Printf("%6d", edges[i][j])
			}
			fmt.Print("\n")
		}

	*/
	fmt.Println(value)
}
