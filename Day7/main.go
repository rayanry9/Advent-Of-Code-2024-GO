package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isPossible(nums *[]int, pos int, value int) bool {
	if value == (*nums)[0] && pos == len(*nums) {
		return true
	}
	if pos >= len(*nums) {
		return false
	}
	return isPossible(nums, pos+1, value+(*nums)[pos]) || isPossible(nums, pos+1, (*nums)[pos]*value)
}

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)
	var value int
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ": ")
		var nums []int
		splits = slices.Concat(splits[0:1], strings.Split(splits[1], " "))
		for i := 0; i < len(splits); i++ {
			val, _ := strconv.Atoi((splits)[i])
			nums = append(nums, val)
		}
		if isPossible(&nums, 1, 0) {
			value += nums[0]
		}
	}
	fmt.Println(value)
}
