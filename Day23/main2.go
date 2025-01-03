package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var connects map[string]map[string]bool
	var localNetwork []string
	var allConnects [][]string

	connects = make(map[string]map[string]bool)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		{
			_, ok := connects[line[0]]
			if !ok {
				connects[line[0]] = make(map[string]bool)
			}
			connects[line[0]][line[1]] = true
		}

		if !slices.Contains(localNetwork, line[0]) {
			localNetwork = append(localNetwork, line[0])
		}
		if !slices.Contains(localNetwork, line[1]) {
			localNetwork = append(localNetwork, line[1])
		}

		allConnects = append(allConnects, line)
	}

	shouldContinue := true

	for shouldContinue {

		shouldContinue = false
		for idx, _ := range allConnects {

			for _, network := range localNetwork {
				shouldBeIn := true

				for _, eachConnect := range allConnects[idx] {
					if network == eachConnect {
						continue
					}
					if !(connects[eachConnect][network] || connects[network][eachConnect]) {
						shouldBeIn = false
					}
				}

				if (!slices.Contains(allConnects[idx], network)) && shouldBeIn {
					allConnects[idx] = append(allConnects[idx], network)
					shouldContinue = true
				}
			}
		}

	}

	var max, maxIdx int
	for idx, val := range allConnects {
		if len(val) > max {
			max = len(val)
			maxIdx = idx
		}
	}

	fmt.Println(len(allConnects[maxIdx]))

	slices.Sort(allConnects[maxIdx])
	for _, val := range allConnects[maxIdx] {
		fmt.Print(val, ",")
	}

}
