package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getSecretNumberAt(num int, times int) int {
	if times == 0 {
		return num
	}
	num = num ^ (num * 64)
	num %= 16777216

	num = num ^ (num / 32)
	num %= 16777216

	num = num ^ (num * 2048)
	num %= 16777216

	return getSecretNumberAt(num, times-1)
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var totalSum int

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		totalSum += getSecretNumberAt(num, 2000)
	}
	fmt.Println(totalSum)

}
