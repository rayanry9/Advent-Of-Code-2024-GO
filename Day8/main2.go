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
								antennas[a][b] = 1
								antennas[i][j] = 1

								vertDist := a - i
								horDist := b - j
								vertDistOri := a - i
								horDistOri := b - j
								for {
									if i-vertDist < 0 || j-horDist < 0 || j-horDist >= SIZE {
										break
									}
									antennas[i-vertDist][j-horDist] = 1
									vertDist += vertDistOri
									horDist += horDistOri
								}
								vertDist = vertDistOri
								horDist = horDistOri
								for {
									if a+vertDist >= SIZE || b+horDist < 0 || b+horDist >= SIZE {
										break
									}
									antennas[a+vertDist][b+horDist] = 1
									vertDist += vertDistOri
									horDist += horDistOri
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
