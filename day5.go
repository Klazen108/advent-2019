package main

import (
	"strconv"
)

func Challenge5_1(program string, input int64) string {
	memory := CsvToIntSlice(program)
	output := RunProgram5(&memory, input)
	return strconv.Itoa(int(output))
}

func RunProgram5(opcodes *[]int64, input int64) int64 {
	result := int64(0)
	pc := 0
	for {
		opcode := Opcode((*opcodes)[pc])
		switch opcode.Code() {
		case 1: //add
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			targetAddr := (*opcodes)[pc+3]

			res := operand1 + operand2
			(*opcodes)[targetAddr] = res
			pc += opcode.ParmCount()
		case 2: //multiply
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			targetAddr := (*opcodes)[pc+3]

			res := operand1 * operand2
			(*opcodes)[targetAddr] = res
			pc += opcode.ParmCount()
		case 3: //input
			targetAddr := (*opcodes)[pc+1]
			(*opcodes)[targetAddr] = input
			pc += opcode.ParmCount()
		case 4: //output
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)

			result = operand1
			pc += opcode.ParmCount()
		case 5:
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			if operand1 != 0 {
				pc = int(operand2)
			} else {
				pc += opcode.ParmCount()
			}
		case 6:
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			if operand1 == 0 {
				pc = int(operand2)
			} else {
				pc += opcode.ParmCount()
			}
		case 7:
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			targetAddr := (*opcodes)[pc+3]

			if operand1 < operand2 {
				(*opcodes)[targetAddr] = 1
			} else {
				(*opcodes)[targetAddr] = 0
			}

			pc += opcode.ParmCount()
		case 8:
			operand1 := GetOperandVal(opcode, opcodes, pc, 0)
			operand2 := GetOperandVal(opcode, opcodes, pc, 1)
			targetAddr := (*opcodes)[pc+3]

			if operand1 == operand2 {
				(*opcodes)[targetAddr] = 1
			} else {
				(*opcodes)[targetAddr] = 0
			}

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

func GetOperandVal(op Opcode, memory *[]int64, pc int, parmIndex int) int64 {
	var operand1 int64
	if op.Mode(parmIndex) == Immediate {
		operand1 = (*memory)[pc+1+parmIndex]
	} else {
		operand1Addr := (*memory)[pc+1+parmIndex]
		operand1 = (*memory)[operand1Addr]
	}
	return operand1
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
	case 5:
		return 3
	case 6:
		return 3
	case 7:
		return 4
	case 8:
		return 4
	case 99:
		return 1
	}
	return 1
}
