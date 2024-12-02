package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("I1.txt")
	if err != nil {
		fmt.Print("COULDNT Open file for reading")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		num := strings.Split(line, " ")
		fmt.Printf("|| %v ||\n", num)
		isIncreasing := 1
		v1, _ := strconv.Atoi(num[0])
		v2, _ := strconv.Atoi(num[1])

		if v1 > v2 {
			isIncreasing = 0
		}

		isTrue := 2

		for i := 0; i < len(num)-1; i++ {
			v1, _ = strconv.Atoi(num[i])
			v2, _ = strconv.Atoi(num[i+1])
			fmt.Printf("%d ", v1)

			if isIncreasing == 1 {
				if !((v2-v1) < 4 && (v2-v1) > 0) {
					isTrue--
					if isTrue == 0 {
						break
					}
				}
			} else if isIncreasing == 0 {
				if !((v1-v2) < 4 && (v1-v2) > 0) {
					isTrue--
					if isTrue == 0 {
						break
					}
				}
			}
		}

		if isTrue > 0 {
			count++
		}
	}

	fmt.Printf("%d", count)
}
