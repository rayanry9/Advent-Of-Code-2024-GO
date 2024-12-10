package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	var arr []int
	backCounter := -1
	for i := 0; i < len(data); i++ {
		if data[i] != 10 {
			if i%2 == 0 {
				backCounter++
			}
			arr = append(arr, int(data[i]-'0'))
		}
	}
	//	fmt.Println(backCounter)
	var value int
	var posValue int
	frontCounter := 0

	for i := 0; i < len(arr) && frontCounter <= backCounter; {
		//fmt.Println(arr)
		if arr[i] == 0 {
			if i%2 == 0 {
				frontCounter++
			}
			i++
			continue
		}
		//fmt.Printf("I: %v | BC : %v | FC : %v | Val : %v\n", i, backCounter, frontCounter, posValue)
		if i%2 == 0 {
			value += posValue * frontCounter
		} else {

			if arr[backCounter*2] == 0 {
				backCounter--
				continue
			}
			value += posValue * backCounter
			arr[backCounter*2]--
		}
		arr[i]--
		posValue++

	}

	fmt.Println(value)
}
