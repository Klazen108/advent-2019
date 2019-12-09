package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/klazen108/advent-2019/intcode"
)

func Challenge7_1(amplifierProgram string) string {
	possiblePhaseSequences := permutations([]int{0, 1, 2, 3, 4})
	maxSignal := 0
	for _, possSeq := range possiblePhaseSequences {
		phaseSequence := make([]string, len(possSeq))
		for i, char := range possSeq {
			phaseSequence[i] = strconv.Itoa(char)
		}
		signal := TestPhaseSequence(amplifierProgram, strings.Join(phaseSequence, ","))
		if signal > maxSignal {
			maxSignal = signal
		}
		fmt.Printf("%s: %d\n", strings.Join(phaseSequence, ","), signal)
	}

	return strconv.Itoa(maxSignal)
}

func Challenge7_2(amplifierProgram string) string {
	possiblePhaseSequences := permutations([]int{5, 6, 7, 8, 9})
	maxSignal := 0
	for _, possSeq := range possiblePhaseSequences {
		phaseSequence := make([]string, len(possSeq))
		for i, char := range possSeq {
			phaseSequence[i] = strconv.Itoa(char)
		}
		signal := TestPhaseSequence(amplifierProgram, strings.Join(phaseSequence, ","))
		if signal > maxSignal {
			maxSignal = signal
		}
		fmt.Printf("%s: %d\n", strings.Join(phaseSequence, ","), signal)
	}

	return strconv.Itoa(maxSignal)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func TestPhaseSequence(amplifierProgram string, phaseSequence string) int {
	program := ParseProgram(amplifierProgram)
	phases := ParseProgram(phaseSequence)
	amplifierChain := make([]*Amplifier, len(phases))

	for i := 0; i < len(phases); i++ {
		amp := NewAmplifier(program, int(phases[i]))
		amplifierChain[i] = &amp
	}

	output := 0
	eee := 0
	pendingInput := true
	for pendingInput {
		pendingInput = false
		for _, amplifier := range amplifierChain {
			curOutput, curPendingInput := amplifier.Compute(output)
			pendingInput = curPendingInput
			output = curOutput
		}
		eee++
		if eee >= 5000 {
			fmt.Printf("wtf")
			break
		}
	}
	return int(output)
}

func (amp *Amplifier) Compute(input int) (int, bool) {
	if !amp.hasRun || amp.computer.IsPendingInput() {
		amp.computer.ProvideInput(intcode.Byte(input))
		amp.computer.Execute()
		amp.hasRun = true

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		amp.output = int(amp.computer.GetOutput())
	}
	return amp.output, amp.computer.IsPendingInput()
}

func NewAmplifier(program []intcode.Byte, phase int) Amplifier {
	comp := intcode.NewComputer(program)
	amp := Amplifier{
		&comp,
		phase,
		false,
		0,
	}
	amp.computer.Reset()
	amp.computer.ProvideInput(intcode.Byte(amp.phase))
	return amp
}

type Amplifier struct {
	computer *intcode.IntcodeComputer
	phase    int
	hasRun   bool
	output   int
}

func ParseProgram(csvInput string) []intcode.Byte {
	csvSlice := strings.Split(csvInput, ",")
	intSlice := make([]intcode.Byte, len(csvSlice))
	for i, elem := range csvSlice {
		intval, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid value in input: '%s'\n", elem)
			os.Exit(103)
		}
		intSlice[i] = intcode.Byte(intval)
	}
	return intSlice
}
