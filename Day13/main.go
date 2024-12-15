package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	counter := 0
	var aX, aY, bX, bY, finalX, finalY int
	var value int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			counter = 0
			continue
		}

		spaceSep := strings.Split(line, " ")
		if counter == 0 {
			aY, _ = strconv.Atoi(spaceSep[3][2:])
			spaceSep = strings.Split(spaceSep[2], ",")
			aX, _ = strconv.Atoi(spaceSep[0][2:])
		} else if counter == 1 {
			bY, _ = strconv.Atoi(spaceSep[3][2:])
			spaceSep = strings.Split(spaceSep[2], ",")
			bX, _ = strconv.Atoi(spaceSep[0][2:])
		} else if counter == 2 {
			finalY, _ = strconv.Atoi(spaceSep[2][2:])
			spaceSep = strings.Split(spaceSep[1], ",")
			finalX, _ = strconv.Atoi(spaceSep[0][2:])

			// Implementation

			d := (aX * bY) - (bX * aY)
			dA := (finalX * bY) - (finalY * bX)
			dB := (aX * finalY) - (aY * finalX)

			if !(d == 0 && dA != 0 && dB != 0) {
				if dA%d == 0 && dB%d == 0 {
					value += ((dA / d) * 3) + (dB / d)
					fmt.Println(dA/d, dB/d)
				}
			}
		}
		counter++
	}
	fmt.Println(value)
}
