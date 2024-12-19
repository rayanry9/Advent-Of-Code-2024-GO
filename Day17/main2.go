package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
3 0
5 5 - > OUT B%8
1 6 - > B = B XOR 6
0 3 - > A = A / 8
4 5 - > B = B XOR C
7 5 - > C = A / 2^B
1 5 - > B = B XOR 5
2 4 - > B = A % 8

*/

func traceBack(regA int, program *[]int, idx int) (bool, int) {
	for i := regA * 8; i < (regA+1)*8; i++ {

		b := i % 8
		b ^= 5
		c := i / (1 << b)
		b ^= c
		b ^= 6
		b %= 8

		if b == (*program)[idx] {
			if idx == 0 {
				return true, i
			}
			is, val := traceBack(i, program, idx-1)
			if is {
				return is, val
			}
		}

	}
	return false, 0
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var regA, lineNumber int
	var program []int

	for scanner.Scan() {
		line := scanner.Text()
		spaceSplit := strings.Split(line, " ")

		if lineNumber == 0 {
			regA, _ = strconv.Atoi(spaceSplit[2])
		} else if lineNumber == 4 {
			commaSplit := strings.Split(spaceSplit[1], ",")
			for _, val := range commaSplit {
				val, _ := strconv.Atoi(val)
				program = append(program, val)
			}
		}
		lineNumber++
	}

	_, regA = traceBack(0, &program, len(program)-1)
	fmt.Println(regA)

}
