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
	if pos >= len(*nums) || value > (*nums)[0] {
		return false
	}
	conc := strconv.Itoa(value) + strconv.Itoa((*nums)[pos])
	concInt, _ := strconv.Atoi(conc)
	//fmt.Printf("%v %v %v\n", value, (*nums)[pos], concInt)
	return isPossible(nums, pos+1, value+(*nums)[pos]) || isPossible(nums, pos+1, (*nums)[pos]*value) || isPossible(nums, pos+1, concInt)
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
		if isPossible(&nums, 2, nums[1]) {
			value += nums[0]
		}
	}
	fmt.Println(value)
}
