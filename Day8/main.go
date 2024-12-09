package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 50

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	var arr [SIZE][SIZE]int
	var antennas [SIZE][SIZE]int

	var count int
	lineNumber := 0
	var max int
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] != '.' {
				arr[lineNumber][i] = int(line[i])
				if int(line[i]) > max {
					max = int(line[i])
				}
			}
		}
		lineNumber++
	}

	for {
		var nextMax int
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				if arr[i][j] > nextMax && arr[i][j] != max {
					nextMax = arr[i][j]
				}

				if arr[i][j] == max {
					for a := i; a < SIZE; a++ {
						for b := 0; b < SIZE; b++ {
							if arr[a][b] == max && (j != b && i != a) {
								if (2*i)-a >= 0 && (2*i)-a < SIZE && (2*j)-b >= 0 && (2*j)-b < SIZE {
									antennas[(2*i)-a][(2*j)-b] = 1
								}
								if (2*a)-i >= 0 && (2*a)-i < SIZE && (2*b)-j >= 0 && (2*b)-j < SIZE {
									antennas[(2*a)-i][(2*b)-j] = 1
								}
							}
						}

					}
					arr[i][j] = 0
				}

			}
		}

		max = nextMax
		if max == 0 {
			break
		}
	}

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if antennas[i][j] > 0 {
				count++
			}
		}
	}
	fmt.Println(count)

}
