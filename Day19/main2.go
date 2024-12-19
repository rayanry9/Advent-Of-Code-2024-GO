package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countPossible(seq string, possibilites *map[string]bool, memo *map[string]int) int {
	if seq == "" {
		return 1
	}
	val, ok := (*memo)[seq]
	if ok {
		return val
	}
	var count int
	for i := range len(seq) {
		_, ok := (*possibilites)[seq[:i+1]]
		if ok {
			count += countPossible(seq[i+1:], possibilites, memo)
		}
	}
	(*memo)[seq] = count
	return count
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	possibilites := make(map[string]bool)
	memo := make(map[string]int)

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
	lineNumber := 0
	for scanner.Scan() {
		seq := scanner.Text()
		count += countPossible(seq, &possibilites, &memo)
		lineNumber++
	}
	fmt.Println(count)

}
