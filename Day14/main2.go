package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const WIDTH = 101
	const HEIGHT = 103

	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	posVel := make([][4]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		elm := strings.Split(line, ",")
		px, _ := strconv.Atoi(strings.Split(elm[0], "=")[1])
		py, _ := strconv.Atoi(strings.Split(elm[1], " ")[0])
		vx, _ := strconv.Atoi(strings.Split(elm[1], "=")[1])
		vy, _ := strconv.Atoi(elm[2])
		posVel = append(posVel, [4]int{px, py, vx, vy})
	}

	for timeAfter := 1; ; timeAfter++ {
		for idx, _ := range posVel {
			posVel[idx][0] += posVel[idx][2]
			if posVel[idx][0] < 0 {
				posVel[idx][0] += WIDTH
			}
			posVel[idx][0] %= WIDTH

			posVel[idx][1] += posVel[idx][3]
			if posVel[idx][1] < 0 {
				posVel[idx][1] += HEIGHT
			}
			posVel[idx][1] %= HEIGHT
		}

		var arr [HEIGHT][WIDTH]int
		for _, val := range posVel {
			arr[val[1]][val[0]]++
		}

		isTrue := false

		heightOfStem := 9

		for i := range WIDTH {
			for j := 0; j < HEIGHT-heightOfStem; j++ {
				exists := true
				for k := 0; k < heightOfStem; k++ {
					if arr[j+k][i] == 0 {
						exists = false
					}
				}
				if exists {
					isTrue = true
					break
				}
			}
			if isTrue {
				break
			}
		}

		if isTrue {
			for i := range HEIGHT {
				for j := range WIDTH {
					if arr[i][j] != 0 {
						fmt.Print(".")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println("")
			}
			fmt.Println(timeAfter)
			break
		}

	}

}
