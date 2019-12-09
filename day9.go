package main

import (
	"strconv"
	"strings"

	"github.com/klazen108/advent-2019/intcode"
)

func Challenge9_1(input string) string {
	program := ParseProgram(input)

	comp := intcode.NewComputer(program)
	comp.Reset()
	comp.ProvideInput(1)
	comp.SetDebug(true)
	comp.Execute()

	outStrs := make([]string, len(comp.InspectOutput()))
	for i, out := range comp.InspectOutput() {
		outStrs[i] = strconv.Itoa(int(out))
	}
	output := strings.Join(outStrs, ",")
	return output
}
