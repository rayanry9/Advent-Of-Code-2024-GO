package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var regA, regB, regC, lineNumber int
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
	for i := 0; i < len(program); i += 2 {
		opCode := program[i]
		operand := program[i+1]
		comboOperand := operand

		switch operand {
		case 4:
			comboOperand = regA
		case 5:
			comboOperand = regB
		case 6:
			comboOperand = regC
		}

		switch opCode {
		case 0:
			regA = regA / (1 << comboOperand)
		case 1:
			regB ^= operand
		case 2:
			regB = comboOperand % 8
		case 3:
			if regA != 0 {
				i = operand
				i -= 2
			}
		case 4:
			regB ^= regC
		case 5:
			fmt.Printf("%d,", comboOperand%8)
		case 6:
			regB = regA / (1 << comboOperand)
		case 7:
			regC = regA / (1 << comboOperand)
		}

	}
	fmt.Print("\n")
	fmt.Println(program)
}
