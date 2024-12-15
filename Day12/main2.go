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
	var vertices [SIZE][SIZE]int

	for i := range SIZE {
		for j := range SIZE {

			// Upper Left Vertice
			if i == 0 && j == 0 {
				vertices[i][j]++
			} else if i == 0 {
				if arr[i][j] != arr[i][j-1] {
					vertices[i][j]++
				}
			} else if j == 0 {
				if arr[i][j] != arr[i-1][j] {
					vertices[i][j]++
				}
			} else {
				if arr[i][j] != arr[i-1][j] && arr[i][j] != arr[i][j-1] {
					vertices[i][j]++
				} else if arr[i][j] == arr[i-1][j] && arr[i][j] == arr[i][j-1] && arr[i][j] != arr[i-1][j-1] {
					vertices[i][j]++
				}
			}

			// Upper Right Vertice
			if i == 0 && j == SIZE-1 {
				vertices[i][j]++
			} else if i == 0 {
				if arr[i][j] != arr[i][j+1] {
					vertices[i][j]++
				}
			} else if j == SIZE-1 {
				if arr[i][j] != arr[i-1][j] {
					vertices[i][j]++
				}
			} else {
				if arr[i][j] != arr[i-1][j] && arr[i][j] != arr[i][j+1] {
					vertices[i][j]++
				} else if arr[i][j] == arr[i-1][j] && arr[i][j] == arr[i][j+1] && arr[i][j] != arr[i-1][j+1] {
					vertices[i][j]++
				}
			}

			// Lower Right Vertice
			if i == SIZE-1 && j == SIZE-1 {
				vertices[i][j]++
			} else if i == SIZE-1 {
				if arr[i][j] != arr[i][j+1] {
					vertices[i][j]++
				}
			} else if j == SIZE-1 {
				if arr[i][j] != arr[i+1][j] {
					vertices[i][j]++
				}
			} else {
				if arr[i][j] != arr[i+1][j] && arr[i][j] != arr[i][j+1] {
					vertices[i][j]++
				} else if arr[i][j] == arr[i+1][j] && arr[i][j] == arr[i][j+1] && arr[i][j] != arr[i+1][j+1] {
					vertices[i][j]++
				}
			}

			// Lower Left Vertice
			if i == SIZE-1 && j == 0 {
				vertices[i][j]++
			} else if i == SIZE-1 {
				if arr[i][j] != arr[i][j-1] {
					vertices[i][j]++
				}
			} else if j == 0 {
				if arr[i][j] != arr[i+1][j] {
					vertices[i][j]++
				}
			} else {
				if arr[i][j] != arr[i+1][j] && arr[i][j] != arr[i][j-1] {
					vertices[i][j]++
				} else if arr[i][j] == arr[i+1][j] && arr[i][j] == arr[i][j-1] && arr[i][j] != arr[i+1][j-1] {
					vertices[i][j]++
				}
			}

			parent[getVal(i, j)] = getVal(i, j)
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

	var sides [SIZE * SIZE]int
	var area [SIZE * SIZE]int
	counters := make([]int, 0)

	for i := range SIZE {
		for j := range SIZE {
			area[findSet(&parent, getVal(i, j))]++
			sides[findSet(&parent, getVal(i, j))] += vertices[i][j]

			if !slices.Contains(counters, findSet(&parent, getVal(i, j))) {
				counters = append(counters, findSet(&parent, getVal(i, j)))
			}
		}
	}
	for _, val := range counters {
		value += sides[val] * area[val]
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
