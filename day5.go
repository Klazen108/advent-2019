package main

import (
	"strconv"
)

func Challenge5_1(program string) string {
	memory := CsvToIntSlice(program)
	output := RunProgram5(&memory)
	return strconv.Itoa(int(output))
}

func RunProgram5(opcodes *[]int64) int64 {
	result := int64(0)
	pc := 0
	for {
		opcode := Opcode((*opcodes)[pc])
		switch opcode.Code() {
		case 1: //add
			var operand1 int64
			if opcode.Mode(0) == Immediate {
				operand1 = (*opcodes)[pc+1]
			} else {
				operand1Addr := (*opcodes)[pc+1]
				operand1 = (*opcodes)[operand1Addr]
			}
			var operand2 int64
			if opcode.Mode(1) == Immediate {
				operand2 = (*opcodes)[pc+2]
			} else {
				operand2Addr := (*opcodes)[pc+2]
				operand2 = (*opcodes)[operand2Addr]
			}

			targetAddr := (*opcodes)[pc+3]

			res := operand1 + operand2
			(*opcodes)[targetAddr] = res
			pc += opcode.ParmCount()
		case 2: //multiply
			var operand1 int64
			if opcode.Mode(0) == Immediate {
				operand1 = (*opcodes)[pc+1]
			} else {
				operand1Addr := (*opcodes)[pc+1]
				operand1 = (*opcodes)[operand1Addr]
			}
			var operand2 int64
			if opcode.Mode(1) == Immediate {
				operand2 = (*opcodes)[pc+2]
			} else {
				operand2Addr := (*opcodes)[pc+2]
				operand2 = (*opcodes)[operand2Addr]
			}

			targetAddr := (*opcodes)[pc+3]

			res := operand1 * operand2
			(*opcodes)[targetAddr] = res
			pc += opcode.ParmCount()
		case 3: //input
			input := int64(1)
			targetAddr := (*opcodes)[pc+1]
			(*opcodes)[targetAddr] = input
			pc += opcode.ParmCount()
		case 4: //output
			var operand1 int64
			if opcode.Mode(0) == Immediate {
				operand1 = (*opcodes)[pc+1]
			} else {
				operand1Addr := (*opcodes)[pc+1]
				operand1 = (*opcodes)[operand1Addr]
			}

			result = operand1
			pc += opcode.ParmCount()
		case 99:
			return result
			pc += opcode.ParmCount()
		default:
			return -1
		}
	}
	return result
}

type Opcode int64

func (o Opcode) Code() int64 {
	opStr := strconv.Itoa(int(o))
	if len(opStr) == 1 {
		return int64(o)
	}
	codeStr := string(opStr[len(opStr)-2:])
	code, _ := strconv.Atoi(codeStr)
	return int64(code)
}

func (o Opcode) Mode(pos int) Mode {
	codeStr := strconv.Itoa(int(o))
	i := len(codeStr) - 1 - 2 - pos
	if i < 0 {
		return Position //value omitted, assume 0
	}
	modeChar := string(codeStr[i])
	mode, _ := strconv.Atoi(modeChar)
	return Mode(mode)
}

type Mode int

const (
	Position  Mode = 0
	Immediate Mode = 1
)

func (o Opcode) ParmCount() int {
	switch o.Code() {
	case 1:
		return 4
	case 2:
		return 4
	case 3:
		return 2
	case 4:
		return 2
	case 99:
		return 1
	}
	return 1
}
