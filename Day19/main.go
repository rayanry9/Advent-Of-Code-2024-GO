package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPossible(seq string, possibilites *map[string]bool) bool {
	if seq == "" {
		return true
	}
	for i := range len(seq) {
		_, ok := (*possibilites)[seq[:i+1]]
		if ok {
			if isPossible(seq[i+1:], possibilites) {
				return true
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	possibilites := make(map[string]bool)

	{
		scanner.Scan()
		line := scanner.Text()
		splits := strings.Split(line, ", ")
		for _, val := range splits {
			possibilites[val] = true
		}
		scanner.Scan()
	}

	var count int
	for scanner.Scan() {
		seq := scanner.Text()
		if isPossible(seq, &possibilites) {
			count++
		}
	}
	fmt.Println(count)

}
