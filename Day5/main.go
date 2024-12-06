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

	var arr [100][100]int
	lineCount := 1
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineCount <= 1176 {
			nums := strings.Split(line, "|")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			arr[num1][num2] = 1
			lineCount++
		} else {
			numStr := strings.Split(line, ",")
			if lineCount == 1177 {
				lineCount++
				continue
			}
			var nums []int
			for j := 0; j < len(numStr); j++ {
				num, _ := strconv.Atoi(numStr[j])
				nums = append(nums, num)
			}
			isTrue := true
			for i := 0; i < len(nums); i++ {
				for j := i + 1; j < len(nums); j++ {
					if arr[nums[j]][nums[i]] == 1 {
						isTrue = false
						break
					}
				}
				if !isTrue {
					break
				}
			}
			if isTrue {
				count += nums[len(nums)/2]
			}
		}
	}
	fmt.Println(count)
}
