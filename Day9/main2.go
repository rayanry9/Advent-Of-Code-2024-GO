package main

import (
	"fmt"
	"os"
	"slices"
)

type memor struct {
	typ  int
	data int
	id   int
}

func main() {
	data, _ := os.ReadFile("./input.txt")
	var arr []memor
	id := 0
	for i := 0; i < len(data); i++ {
		if data[i] != 10 {
			if i%2 == 0 {
				arr = append(arr, memor{1, int(data[i] - '0'), id})
				id++
			} else {
				arr = append(arr, memor{0, int(data[i] - '0'), 0})
			}
		}
	}
	var value int

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i].typ == 1 {
			for j := 0; j < i; j++ {
				if arr[j].typ == 0 && arr[j].data >= arr[i].data {
					arr[j].data -= arr[i].data
					arr = slices.Concat(arr[:j], arr[i:i+1], arr[j:])
					i++
					arr[i].typ = 0
					arr[i].id = 0
					break
				}
			}
		}
	}

	posValue := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i].data; j++ {
			value += posValue * arr[i].id
			posValue++
		}
	}

	fmt.Println(value)
}
