package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var count int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "X"[0] {

				// Down Right Diagonal

				if i+3 < len(lines) && j+3 < len(lines[i]) {
					if lines[i+1][j+1] == "M"[0] {
						if lines[i+2][j+2] == "A"[0] {
							if lines[i+3][j+3] == "S"[0] {
								count++
							}
						}
					}
				}

				// Down
				if i+3 < len(lines) {
					if lines[i+1][j] == "M"[0] {
						if lines[i+2][j] == "A"[0] {
							if lines[i+3][j] == "S"[0] {
								count++
							}
						}
					}
				}

				// Right

				if j+3 < len(lines[i]) {
					if lines[i][j+1] == "M"[0] {
						if lines[i][j+2] == "A"[0] {
							if lines[i][j+3] == "S"[0] {
								count++
							}
						}
					}
				}

				// Down Left

				if i+3 < len(lines) && j-3 >= 0 {
					if lines[i+1][j-1] == "M"[0] {
						if lines[i+2][j-2] == "A"[0] {
							if lines[i+3][j-3] == "S"[0] {
								count++
							}
						}
					}
				}

				// Left

				if j-3 >= 0 {
					if lines[i][j-1] == "M"[0] {
						if lines[i][j-2] == "A"[0] {
							if lines[i][j-3] == "S"[0] {
								count++
							}
						}
					}
				}

				// Up Left

				if i-3 >= 0 && j-3 >= 0 {
					if lines[i-1][j-1] == "M"[0] {
						if lines[i-2][j-2] == "A"[0] {
							if lines[i-3][j-3] == "S"[0] {
								count++
							}
						}
					}
				}

				// Up

				if i-3 >= 0 {
					if lines[i-1][j] == "M"[0] {
						if lines[i-2][j] == "A"[0] {
							if lines[i-3][j] == "S"[0] {
								count++
							}
						}
					}
				}

				// Up right

				if i-3 >= 0 && j+3 < len(lines[i]) {
					if lines[i-1][j+1] == "M"[0] {
						if lines[i-2][j+2] == "A"[0] {
							if lines[i-3][j+3] == "S"[0] {
								count++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)

}
