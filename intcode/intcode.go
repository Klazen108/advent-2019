package intcode

import "fmt"

type Byte int

type IntcodeComputer struct {
	memory []Byte
	pc     int
	halt   bool

	outputBuffer []Byte
	inputBuffer  []Byte
}

func NewComputer(memory []Byte) IntcodeComputer {
	return IntcodeComputer{
		memory,
		0,
		false,
		[]Byte{},
		[]Byte{},
	}
}

func (comp IntcodeComputer) Reset() {
	comp.pc = 0
	comp.outputBuffer = make([]Byte, 0)
	comp.inputBuffer = make([]Byte, 0)
	comp.halt = false
}

func (comp IntcodeComputer) Load(memory []Byte) {
	comp.memory = memory
}

func (comp *IntcodeComputer) Execute() {
	i := 0
	for comp.Tick() {
		i++
		if i > 1000 {
			panic("Infinite loop detected")
		}
	}
}

func (comp *IntcodeComputer) Tick() bool {
	inst := LoadInstruction(comp.pc, comp.memory)
	fmt.Printf("%T\n", inst)
	if inst.Execute(comp) {
		comp.pc += inst.Length()
	}
	return !comp.halt
}

func (comp IntcodeComputer) GetInput() Byte {
	if len(comp.inputBuffer) == 0 {
		panic("Error: Input buffer empty")
	}
	val := comp.inputBuffer[0]
	comp.inputBuffer = comp.inputBuffer[1:]
	return val
}

func (comp *IntcodeComputer) Output(val Byte) {
	fmt.Printf("Outputting: %d\n", val)
	comp.outputBuffer = append(comp.outputBuffer, val)
}

func (comp *IntcodeComputer) ProvideInput(val Byte) {
	comp.inputBuffer = append(comp.inputBuffer, val)
}

func (comp *IntcodeComputer) Halt() {
	comp.halt = true
}
