package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNextSecretNumber(num int) int {
	num = num ^ (num * 64)
	num %= 16777216

	num = num ^ (num / 32)
	num %= 16777216

	num = num ^ (num * 2048)
	num %= 16777216

	return num
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	seqVal := make(map[string]int)

	for scanner.Scan() {
		hasSold := make(map[string]bool)

		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		var fiveNums [5]int
		fiveNums[0] = num
		fiveNums[1] = getNextSecretNumber(num)
		fiveNums[2] = getNextSecretNumber(fiveNums[1])
		fiveNums[3] = getNextSecretNumber(fiveNums[2])
		fiveNums[4] = getNextSecretNumber(fiveNums[3])

		seq := strconv.Itoa((fiveNums[1] % 10) - (fiveNums[0] % 10))
		seq = seq + strconv.Itoa((fiveNums[2]%10)-(fiveNums[1]%10))
		seq = seq + strconv.Itoa((fiveNums[3]%10)-(fiveNums[2]%10))
		seq = seq + strconv.Itoa((fiveNums[4]%10)-(fiveNums[3]%10))
		seqVal[seq] = fiveNums[4] % 10

		for range 1996 {
			fiveNums[0] = fiveNums[1]
			fiveNums[1] = fiveNums[2]
			fiveNums[2] = fiveNums[3]
			fiveNums[3] = fiveNums[4]
			fiveNums[4] = getNextSecretNumber(fiveNums[3])

			seq = strconv.Itoa((fiveNums[1] % 10) - (fiveNums[0] % 10))
			seq = seq + strconv.Itoa((fiveNums[2]%10)-(fiveNums[1]%10))
			seq = seq + strconv.Itoa((fiveNums[3]%10)-(fiveNums[2]%10))
			seq = seq + strconv.Itoa((fiveNums[4]%10)-(fiveNums[3]%10))

			val, ok := seqVal[seq]
			_, sold := hasSold[seq]
			if !sold {
				if ok {
					seqVal[seq] = val + fiveNums[4]%10
				} else {
					seqVal[seq] = fiveNums[4] % 10
				}
				hasSold[seq] = true
			}
		}
	}

	var max int
	for _, val := range seqVal {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)

}
