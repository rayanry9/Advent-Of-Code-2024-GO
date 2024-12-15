package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	for range 25 {
		for idx := len(line) - 1; idx >= 0; idx-- {

			if line[idx] == "0" {
				line[idx] = "1"
			} else if len(line[idx])%2 == 0 {
				num1, _ := strconv.Atoi(line[idx][:len(line[idx])/2])
				num2, _ := strconv.Atoi(line[idx][len(line[idx])/2:])
				numStr1 := strconv.Itoa(num1)
				numStr2 := strconv.Itoa(num2)
				line = slices.Concat(line[:idx], []string{numStr1, numStr2}, line[idx+1:])
			} else {
				num, _ := strconv.Atoi(line[idx])
				num *= 2024
				line[idx] = strconv.Itoa(num)
			}

		}
	}
	fmt.Println(len(line))

}
