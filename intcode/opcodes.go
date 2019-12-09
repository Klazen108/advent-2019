package intcode

import (
	"fmt"
	"strconv"
)

func (op RawInstruction) GetParamVal(parmIndex int, memory []Byte, forceImm bool) Byte {
	var operand1 Byte
	if (op.ParamMode(parmIndex, memory) == Immediate) || forceImm {
		operand1 = memory[op.rawLoc+1+parmIndex]
	} else {
		operand1Addr := memory[op.rawLoc+1+parmIndex]
		operand1 = memory[operand1Addr]
	}
	return operand1
}

func (op RawInstruction) ParamMode(pos int, memory []Byte) Mode {
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

func LoadInstruction(pos int, memory []Byte) Instruction {
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
		return SltInst(rawInst)
	case 8:
		return SeqInst(rawInst)
	case 99:
		return HltInst(rawInst)
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
)

type Instruction interface {
	//if true, increment PC after execution
	Execute(comp *IntcodeComputer) bool
	Length() int
}

type AddInst RawInstruction

func (i AddInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)
	targetAddr := RawInstruction(i).GetParamVal(2, comp.memory, true)
	res := operand1 + operand2
	comp.memory[targetAddr] = res
	return true
}
func (i AddInst) Length() int {
	return 4
}

type MulInst RawInstruction

func (i MulInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)
	targetAddr := RawInstruction(i).GetParamVal(2, comp.memory, true)

	res := operand1 * operand2
	comp.memory[targetAddr] = res
	return true
}
func (i MulInst) Length() int {
	return 4
}

type InpInst RawInstruction

func (i InpInst) Execute(comp *IntcodeComputer) bool {
	targetAddr := RawInstruction(i).GetParamVal(0, comp.memory, true)
	comp.memory[targetAddr] = comp.GetInput()
	return true
}
func (i InpInst) Length() int {
	return 2
}

type OutInst RawInstruction

func (i OutInst) Execute(comp *IntcodeComputer) bool {
	output := RawInstruction(i).GetParamVal(0, comp.memory, false)
	comp.Output(output)
	return true
}
func (i OutInst) Length() int {
	return 2
}

type JneInst RawInstruction

func (i JneInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)

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
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)

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
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)
	targetAddr := RawInstruction(i).GetParamVal(2, comp.memory, true)

	if operand1 < operand2 {
		comp.memory[targetAddr] = 1
	} else {
		comp.memory[targetAddr] = 0
	}
	return true
}
func (i SltInst) Length() int {
	return 4
}

type SeqInst RawInstruction

func (i SeqInst) Execute(comp *IntcodeComputer) bool {
	operand1 := RawInstruction(i).GetParamVal(0, comp.memory, false)
	operand2 := RawInstruction(i).GetParamVal(1, comp.memory, false)
	targetAddr := RawInstruction(i).GetParamVal(2, comp.memory, true)

	if operand1 == operand2 {
		comp.memory[targetAddr] = 1
	} else {
		comp.memory[targetAddr] = 0
	}
	return true
}
func (i SeqInst) Length() int {
	return 4
}

type HltInst RawInstruction

func (i HltInst) Execute(comp *IntcodeComputer) bool {
	comp.Halt()
	return false
}
func (i HltInst) Length() int {
	return 1
}
