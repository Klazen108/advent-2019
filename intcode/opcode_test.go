package intcode

import (
	"fmt"
	"testing"
)

func TestLoadInstruction_1(t *testing.T) {
	memory := make(map[int]Byte)
	memory[0] = 99
	inst := LoadInstruction(0, memory)

	_, ok := inst.(HltInst)

	if !ok {
		t.Errorf("Expected a halt instruction")
	}
}

func TestLoadInstruction_4(t *testing.T) {
	memory := make(map[int]Byte)
	memory[0] = 1
	memory[1] = 2
	memory[2] = 3
	memory[3] = 1
	inst := LoadInstruction(0, memory)

	_, ok := inst.(AddInst)

	if !ok {
		t.Errorf("Expected an add instruction")
	}
}

func TestAddInstruction(t *testing.T) {
	comp := NewComputer([]Byte{1101, 2, 3, 1})

	inst := LoadInstruction(0, comp.memory)
	inst.Execute(&comp)

	if comp.memory[1] != 5 {
		t.Errorf("Expected 5 at position 1, got %d", comp.memory[1])
	}
}

func TestInpInstruction(t *testing.T) {
	comp := NewComputer([]Byte{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8})

	inst := LoadInstruction(0, comp.memory)
	comp.ProvideInput(8)
	fmt.Printf("%+v", comp.inputBuffer)
	inst.Execute(&comp)

	if comp.memory[9] != 8 {
		t.Errorf("Expected 8 at position 9, got %d", comp.memory[8])
	}
}

func TestOutInstruction(t *testing.T) {
	comp := NewComputer([]Byte{104, 1337})

	inst := LoadInstruction(0, comp.memory)
	inst.Execute(&comp)
	if len(comp.outputBuffer) == 0 {
		t.Errorf("Expected output")
		return
	}

	if comp.outputBuffer[0] != 1337 {
		t.Errorf("Expected output of 1337, got %d", comp.outputBuffer[0])
	}
}
