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
	timeAfter := 100

	var topLeft, topRight, botLeft, botRight int

	for idx := range posVel {
		posVel[idx][0] += posVel[idx][2] * timeAfter
		posVel[idx][0] %= WIDTH
		posVel[idx][0] += WIDTH
		posVel[idx][0] %= WIDTH
		posVel[idx][1] += posVel[idx][3] * timeAfter
		posVel[idx][1] %= HEIGHT
		posVel[idx][1] += HEIGHT
		posVel[idx][1] %= HEIGHT

		if posVel[idx][0] < WIDTH/2 && posVel[idx][1] < HEIGHT/2 {
			topLeft++
		} else if posVel[idx][0] < WIDTH/2 && posVel[idx][1] > HEIGHT/2 {
			botLeft++
		} else if posVel[idx][0] > WIDTH/2 && posVel[idx][1] < HEIGHT/2 {
			topRight++
		} else if posVel[idx][0] > WIDTH/2 && posVel[idx][1] > HEIGHT/2 {
			botRight++
		}
	}

	var arr [HEIGHT][WIDTH]int
	for _, val := range posVel {
		arr[val[1]][val[0]]++
	}
	for i := range HEIGHT {
		if i == HEIGHT/2 {
			fmt.Println("")
			continue
		}
		for j := range WIDTH {
			if j == WIDTH/2 {
				fmt.Print("  ")
				continue
			}
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println("")
	}

	fmt.Println(topRight, topLeft, botRight, botLeft)
	fmt.Println(topRight * topLeft * botRight * botLeft)

}
