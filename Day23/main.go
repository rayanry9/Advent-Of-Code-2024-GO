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
	var threeConnects [][3]string

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
	}

	for kys, val := range connects {
		for kys2 := range val {
			for _, network := range localNetwork {
				if !(kys[0] == 't' || kys2[0] == 't' || network[0] == 't') {
					continue
				}

				if (connects[network][kys] || connects[kys][network]) && (connects[network][kys2] || connects[kys2][network]) {

					condition := slices.Contains(threeConnects, [3]string{kys, kys2, network})
					condition = condition || slices.Contains(threeConnects, [3]string{kys, network, kys2})
					condition = condition || slices.Contains(threeConnects, [3]string{kys2, kys, network})
					condition = condition || slices.Contains(threeConnects, [3]string{kys2, network, kys})
					condition = condition || slices.Contains(threeConnects, [3]string{network, kys2, kys})
					condition = condition || slices.Contains(threeConnects, [3]string{network, kys, kys2})

					if !condition {
						threeConnects = append(threeConnects, [3]string{kys, kys2, network})
					}
				}
			}
		}
	}

	fmt.Println(len(threeConnects))

}
