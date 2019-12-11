package main

import (
	"strconv"
	"strings"

	"github.com/klazen108/advent-2019/intcode"
)

func Challenge9_1(pgm string, mode int) string {
	program := ParseProgram(pgm)

	comp := intcode.NewComputer(program)
	comp.Reset()
	comp.ProvideInput(intcode.Byte(mode))
	comp.Execute()

	outStrs := make([]string, len(comp.InspectOutput()))
	for i, out := range comp.InspectOutput() {
		outStrs[i] = strconv.Itoa(int(out))
	}
	output := strings.Join(outStrs, ",")
	return output
}
