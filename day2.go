package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge2_1(program string) string {
	memory := CsvToIntSlice(program)
	//fix program
	memory[1] = 12
	memory[2] = 2
	RunProgram(&memory)
	fmt.Fprintf(os.Stderr, "Value at position 0: %d\n", memory[0])
	return strconv.FormatInt(memory[0], 10)
}

func Challenge2_2(program string) string {
	for noun := int64(0); noun <= 99; noun++ {
		for verb := int64(0); verb <= 99; verb++ {
			memory := CsvToIntSlice(program)
			memory[1] = noun
			memory[2] = verb
			RunProgram(&memory)
			if memory[0] == 19690720 {
				fmt.Fprintf(os.Stderr, "Noun: %d Verb: %d\n", memory[1], memory[2])
				return fmt.Sprintf("%02d%02d", memory[1], memory[2])
			}
		}
	}
	return "0000"
}

func RunProgram(opcodes *[]int64) {
	pc := 0
	for {
		opcode := (*opcodes)[pc]
		switch opcode {
		case 1:
			operand1Addr := (*opcodes)[pc+1]
			operand2Addr := (*opcodes)[pc+2]
			targetAddr := (*opcodes)[pc+3]

			operand1 := (*opcodes)[operand1Addr]
			operand2 := (*opcodes)[operand2Addr]

			result := operand1 + operand2
			(*opcodes)[targetAddr] = result
		case 2:
			operand1Addr := (*opcodes)[pc+1]
			operand2Addr := (*opcodes)[pc+2]
			targetAddr := (*opcodes)[pc+3]

			operand1 := (*opcodes)[operand1Addr]
			operand2 := (*opcodes)[operand2Addr]

			result := operand1 * operand2
			(*opcodes)[targetAddr] = result
		case 99:
			return
		}
		pc += 4
	}
}

func CsvToIntSlice(csv string) []int64 {
	csvSlice := strings.Split(csv, ",")
	intSlice := make([]int64, len(csvSlice))
	for i, elem := range csvSlice {
		intval, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid value in input: '%s'\n", elem)
			os.Exit(103)
		}
		intSlice[i] = intval
	}
	return intSlice
}
