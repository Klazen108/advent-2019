package intcode

import (
	"testing"
)

func TestProvideInput(t *testing.T) {
	comp := NewComputer([]Byte{})

	if len(comp.inputBuffer) != 0 {
		t.Errorf("Expected empty input buffer")
	}

	comp.ProvideInput(8)

	if len(comp.inputBuffer) != 1 {
		t.Errorf("Expected one value in input buffer")
	}
}

func TestDay5Example(t *testing.T) {
	comp := NewComputer([]Byte{1002, 4, 3, 4, 33})

	inst := LoadInstruction(0, comp.memory)
	inst.Execute(&comp)

	if comp.memory[4] != 99 {
		t.Errorf("Expected 99 at position 4, got %d", comp.memory[4])
	}
}

func TestDay52Example1Pos(t *testing.T) {
	comp := NewComputer([]Byte{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8})

	comp.Reset()
	comp.ProvideInput(8)
	comp.Execute()

	if len(comp.outputBuffer) == 0 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 1 {
		t.Errorf("Expected output of 1, was %d", comp.outputBuffer[0])
	}
}

func TestDay52Example1Neg(t *testing.T) {
	comp := NewComputer([]Byte{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8})

	comp.Reset()
	comp.ProvideInput(7)
	comp.Execute()

	if len(comp.outputBuffer) == 0 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 0 {
		t.Errorf("Expected output of 0, was %d", comp.outputBuffer[0])
	}
}

func TestDay52Example4Pos(t *testing.T) {
	comp := NewComputer([]Byte{3, 3, 1107, -1, 8, 3, 4, 3, 99})

	comp.Reset()
	comp.ProvideInput(7)
	comp.Execute()

	if len(comp.outputBuffer) == 0 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 1 {
		t.Errorf("Expected output of 1, was %d", comp.outputBuffer[0])
	}
}

func TestDay52Example4Neg(t *testing.T) {
	comp := NewComputer([]Byte{3, 3, 1107, -1, 8, 3, 4, 3, 99})

	comp.Reset()
	comp.ProvideInput(8)
	comp.Execute()

	if len(comp.outputBuffer) == 0 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 0 {
		t.Errorf("Expected output of 0, was %d", comp.outputBuffer[0])
	}
}

func TestDay52Example5Eq(t *testing.T) {
	comp := NewComputer([]Byte{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99})

	comp.Reset()
	comp.ProvideInput(8)
	comp.Execute()

	if len(comp.outputBuffer) == 1000 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 1000 {
		t.Errorf("Expected output of 1000, was %d", comp.outputBuffer[1000])
	}
}

func TestDay52Example5Lt(t *testing.T) {
	comp := NewComputer([]Byte{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99})

	comp.Reset()
	comp.ProvideInput(7)
	comp.Execute()

	if len(comp.outputBuffer) == 999 {
		t.Errorf("Expected output")
		return
	}
	if comp.outputBuffer[0] != 999 {
		t.Errorf("Expected output of 999, was %d", comp.outputBuffer[0])
	}
}
