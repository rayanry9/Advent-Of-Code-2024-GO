package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(reader *bufio.Reader, str string) bool {
	b, err := reader.ReadByte()
	if err != nil {
		os.Exit(0)
	}

	fmt.Print(string(b))
	if b == str[0] {
		return true
	} else {
		return false
	}
}

func isNumber(reader *bufio.Reader) (bool, int) {
	b, err := reader.Peek(1)

	if err != nil {
		os.Exit(0)
	}

	fmt.Print(string(b))
	if b[0]-"0"[0] >= 0 && b[0]-"0"[0] < 10 {
		b, err := reader.ReadByte()
		if err != nil {
			os.Exit(0)
		}
		return true, int(b - "0"[0])
	} else {
		return false, 0
	}
}

func getNumber(reader *bufio.Reader) (bool, int) {
	isNum, num1 := isNumber(reader)

	if isNum {
		isNum, num2 := isNumber(reader)

		if isNum {
			isNum, num3 := isNumber(reader)
			if isNum {
				return true, (num1 * 100) + (num2 * 10) + (num3 * 1)
			} else {
				return true, (num1 * 10) + (num2 * 1)
			}
		} else {
			return true, num1
		}
	} else {
		return false, 0
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	reader := bufio.NewReader(file)

	mult := 0
	for {

		if check(reader, "m") {
			if check(reader, "u") {
				if check(reader, "l") {
					if check(reader, "(") {
						isNum, num1 := getNumber(reader)
						if isNum {
							if check(reader, ",") {
								isNum, num2 := getNumber(reader)
								if isNum {
									if check(reader, ")") {
										mult += num1 * num2
										println(mult)
									}
								}
							}
						}
					}
				}
			}
		}

	}
}
