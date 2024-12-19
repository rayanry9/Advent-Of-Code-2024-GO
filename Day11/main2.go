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
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for idx := len(line) - 1; idx >= 0; idx-- {
		val, _ := map1[line[idx]]
		map1[line[idx]] = val + 1
	}

	times := 75

	for i := range times {
		if i%2 == 0 {
			map2 = make(map[string]int)
			for key, val := range map1 {
				if key == "0" {
					map2["1"] += val
				} else if len(key)%2 == 0 {
					map2[key[:len(key)/2]] += val
					num, _ := strconv.Atoi(key[len(key)/2:])
					map2[strconv.Itoa(num)] += val
				} else {
					num, _ := strconv.Atoi(key)
					num *= 2024
					numStr := strconv.Itoa(num)
					map2[numStr] += val
				}
			}
		} else {

			map1 = make(map[string]int)
			for key, val := range map2 {
				if key == "0" {
					map1["1"] += val
				} else if len(key)%2 == 0 {
					map1[key[:len(key)/2]] += val
					num, _ := strconv.Atoi(key[len(key)/2:])
					map1[strconv.Itoa(num)] += val
				} else {
					num, _ := strconv.Atoi(key)
					num *= 2024
					map1[strconv.Itoa(num)] += val
				}
			}
		}

	}

	totalCount := 0

	if times%2 == 0 {
		for _, val := range map1 {
			totalCount += val
		}
	} else {
		for _, val := range map2 {
			totalCount += val
		}
	}
	fmt.Println(totalCount)

}
