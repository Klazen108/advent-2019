package intcode

import (
	"fmt"
	"strconv"
)

func (op RawInstruction) GetParamVal(parmIndex int, memory map[int]Byte, forceImm bool, relativeBase int) Byte {
	var operand1 Byte
	if (op.ParamMode(parmIndex, memory) == Immediate) || forceImm {
		operand1 = memory[op.rawLoc+1+parmIndex]
		fmt.Printf("IMM: %d\n", operand1)
	} else if op.ParamMode(parmIndex, memory) == Relative {
		offset := memory[op.rawLoc+1+parmIndex]
		//operand1Addr := memory[relativeBase+int(offset)]
		operand1Addr := relativeBase + int(offset)
		operand1 = memory[operand1Addr]
		fmt.Printf("REL: mem[%d + %d]=%d\n", relativeBase, int(offset), operand1)
	} else {
		operand1Addr := int(memory[op.rawLoc+1+parmIndex])
		operand1 = memory[operand1Addr]
		fmt.Printf("POS: mem[%d]=%d\n", operand1Addr, operand1)
	}
	return operand1
}

func (op RawInstruction) GetTargetAddr(parmIndex int, memory map[int]Byte, relativeBase int) int {
	var targetAddr int
	value := int(memory[op.rawLoc+1+parmIndex])
	if op.ParamMode(0, memory) == Relative {
		targetAddr = relativeBase + value
	} else {
		targetAddr = value
	}
	return targetAddr
}

func (op RawInstruction) ParamMode(pos int, memory map[int]Byte) Mode {
	rawOp := memory[op.rawLoc]

	codeStr := strconv.Itoa(int(rawOp))
	i := len(codeStr) - 1 - 2 - pos
	if i < 0 {
		return Position //value omitted, assume 0
	}
	modeChar := string(codeStr[i])
	mode, _ := strconv.Atoi(modeChar)
	return Mode(mode)
}

func LoadInstruction(pos int, memory map[int]Byte) Instruction {
	rawInst := RawInstruction{
		pos,
		make([]Operand, 0),
	}

	opcode := Opcode(memory[pos])
	var inst Instruction
	switch opcode.Code() {
	case 1:
		inst = AddInst(rawInst)
	case 2:
		inst = MulInst(rawInst)
	case 3:
		inst = InpInst(rawInst)
	case 4:
		inst = OutInst(rawInst)
	case 5:
		inst = JneInst(rawInst)
	case 6:
		inst = JeqInst(rawInst)
	case 7:
		inst = SltInst(rawInst)
	case 8:
		inst = SeqInst(rawInst)
	case 9:
		inst = SrbInst(rawInst)
	case 99:
		inst = HltInst(rawInst)
	default:
		panic(fmt.Sprintf("Unsupported Opcode: %d", opcode.Code()))
	}

	for i := 1; i < inst.Length(); i++ {
		rawInst.operands = append(rawInst.operands, Operand{
			memory[pos],
			opcode.Mode(i - 1),
		})
	}

	return inst
}

type Opcode int

func (op Opcode) Code() int {
	opStr := strconv.Itoa(int(op))
	if len(opStr) == 1 {
		return int(op)
	}
	codeStr := string(opStr[len(opStr)-2:])
	code, _ := strconv.Atoi(codeStr)
	return int(code)
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

type RawInstruction struct {
	rawLoc   int
	operands []Operand
}

type Operand struct {
	value    Byte
	addrMode Mode
}

type Mode int

const (
	Position  Mode = 0
	Immediate Mode = 1
	Relative  Mode = 2
)

type Instruction interface {
	//if true, increment PC after execution
	Execute(comp *IntcodeComputer) bool
	Length() int
}

type AddInst RawInstruction

func (i AddInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)
	targetAddr := RawInstruction(i).GetTargetAddr(2, comp.memory, comp.relativeBase)
	res := operand1 + operand2
	comp.memory[int(targetAddr)] = res
	return true
}
func (i AddInst) Length() int {
	return 4
}

type MulInst RawInstruction

func (i MulInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)
	targetAddr := RawInstruction(i).GetTargetAddr(2, comp.memory, comp.relativeBase)

	res := operand1 * operand2
	comp.memory[int(targetAddr)] = res
	return true
}
func (i MulInst) Length() int {
	return 4
}

type InpInst RawInstruction

func (i InpInst) Execute(comp *IntcodeComputer) bool {
	targetAddr := RawInstruction(i).GetTargetAddr(0, comp.memory, comp.relativeBase)
	value, hadInput := comp.GetInput()
	if !hadInput {
		comp.Halt(true)
		return false
	}
	comp.memory[targetAddr] = value
	return true
}
func (i InpInst) Length() int {
	return 2
}

type OutInst RawInstruction

func (i OutInst) Execute(comp *IntcodeComputer) bool {
	output := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	comp.Output(output)
	return true
}
func (i OutInst) Length() int {
	return 2
}

type JneInst RawInstruction

func (i JneInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)

	if operand1 != 0 {
		comp.pc = int(operand2)
		return false
	}
	return true
}
func (i JneInst) Length() int {
	return 3
}

type JeqInst RawInstruction

func (i JeqInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)

	if operand1 == 0 {
		comp.pc = int(operand2)
		return false
	}
	return true
}
func (i JeqInst) Length() int {
	return 3
}

type SltInst RawInstruction

func (i SltInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)
	targetAddr := RawInstruction(i).GetTargetAddr(2, comp.memory, comp.relativeBase)

	if operand1 < operand2 {
		comp.memory[int(targetAddr)] = 1
	} else {
		comp.memory[int(targetAddr)] = 0
	}
	return true
}
func (i SltInst) Length() int {
	return 4
}

type SeqInst RawInstruction

func (i SeqInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false, comp.relativeBase)
	targetAddr := RawInstruction(i).GetTargetAddr(2, comp.memory, comp.relativeBase)

	if operand1 == operand2 {
		comp.memory[int(targetAddr)] = 1
	} else {
		comp.memory[int(targetAddr)] = 0
	}
	return true
}
func (i SeqInst) Length() int {
	return 4
}

type HltInst RawInstruction

func (i HltInst) Execute(comp *IntcodeComputer) bool {
	comp.Halt(false)
	return false
}
func (i HltInst) Length() int {
	return 1
}

type SrbInst RawInstruction

func (i SrbInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false, comp.relativeBase)
	if comp.debug {
		fmt.Printf("relBase addend = %d\n", operand1)
	}

	comp.relativeBase += int(operand1)
	if comp.debug {
		fmt.Printf("relBase now = %d\n", comp.relativeBase)
	}
	return true
}
func (i SrbInst) Length() int {
	return 2
}
