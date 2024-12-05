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
			dlu := false
			dld := false
			if lines[i][j] == "A"[0] {
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(lines) && j+1 < len(lines[i]) {
					// Diagonal Left Up

					if lines[i-1][j-1] == "M"[0] {
						if lines[i+1][j+1] == "S"[0] {
							dlu = true
						}
					} else if lines[i-1][j-1] == "S"[0] {
						if lines[i+1][j+1] == "M"[0] {
							dlu = true
						}
					}

					// Diagonal Left Down

					if lines[i+1][j-1] == "M"[0] {
						if lines[i-1][j+1] == "S"[0] {
							dld = true
						}
					} else if lines[i+1][j-1] == "S"[0] {
						if lines[i-1][j+1] == "M"[0] {
							dld = true
						}
					}

					if dlu && dld {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)

}
