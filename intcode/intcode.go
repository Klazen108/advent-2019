package intcode

import "fmt"

type Byte int

type IntcodeComputer struct {
	memory       map[int]Byte
	pc           int
	halt         bool
	pendingInput bool
	relativeBase int
	debug        bool

	outputBuffer []Byte
	inputBuffer  []Byte
}

func NewComputer(memory []Byte) IntcodeComputer {
	mem := make(map[int]Byte)
	for i, b := range memory {
		mem[i] = b
	}
	return IntcodeComputer{
		mem,
		0,
		false,
		false,
		0,
		false,
		[]Byte{},
		[]Byte{},
	}
}

func (comp *IntcodeComputer) SetDebug(debugMode bool) {
	comp.debug = debugMode
}

func (comp *IntcodeComputer) Reset() {
	comp.pc = 0
	comp.outputBuffer = make([]Byte, 0)
	comp.inputBuffer = make([]Byte, 0)
	comp.halt = false
}

func (comp *IntcodeComputer) Load(memory []Byte) {
	comp.memory = make(map[int]Byte)
	for i, b := range memory {
		comp.memory[i] = b
	}
}

func (comp *IntcodeComputer) Execute() {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf(
				"Kernel Panic, opcode [%T] [%d,%d,%d,%d] pc [%d] rb [%d] error [%v]",
				LoadInstruction(comp.pc, comp.memory),
				comp.memory[comp.pc],
				comp.memory[comp.pc+1],
				comp.memory[comp.pc+2],
				comp.memory[comp.pc+3],
				comp.pc,
				comp.relativeBase,
				r))
		}
	}()

	comp.halt = false
	i := 0
	for comp.Tick() {
		i++
		if i > 1000000 {
			panic("Infinite loop detected")
		}
	}
	//fmt.Printf("PC=%d\n", comp.pc)
	//fmt.Printf("Execution Complete\n")
}

func (comp *IntcodeComputer) Tick() bool {
	inst := LoadInstruction(comp.pc, comp.memory)
	if comp.debug {
		fmt.Printf("%T\n", inst)
	}
	if inst.Execute(comp) {
		comp.pc += inst.Length()
	}
	//fmt.Printf("PC=%d\n", comp.pc)
	return !comp.halt
}

// GetInput gets an input from the computer.
// If returns true, then a value may be processed. If false, computer
// should halt and wait for input - reexecution will proceed
// from the halted instruction
func (comp *IntcodeComputer) GetInput() (Byte, bool) {
	if len(comp.inputBuffer) == 0 {
		return 0, false
	}
	val := comp.inputBuffer[0]
	comp.inputBuffer = comp.inputBuffer[1:]
	return val, true
}

func (comp *IntcodeComputer) Output(val Byte) {
	if comp.debug {
		fmt.Printf(">>out:%d\n", val)
	}
	comp.outputBuffer = append(comp.outputBuffer, val)
}

func (comp *IntcodeComputer) ProvideInput(val Byte) {
	comp.inputBuffer = append(comp.inputBuffer, val)
}

func (comp *IntcodeComputer) Halt(pendingInput bool) {
	comp.pendingInput = pendingInput
	comp.halt = true
}

func (comp *IntcodeComputer) GetOutput() Byte {
	if len(comp.outputBuffer) == 0 {
		panic("Error: Output buffer empty")
	}
	val := comp.outputBuffer[0]
	comp.outputBuffer = comp.outputBuffer[1:]
	return val
}

func (comp *IntcodeComputer) InspectOutput() []Byte {
	return comp.outputBuffer
}

func (comp *IntcodeComputer) IsPendingInput() bool {
	return comp.pendingInput
}

func (comp *IntcodeComputer) GetPC() int {
	return comp.pc
}
