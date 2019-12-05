package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestDay2_1(t *testing.T) {
	result := Challenge2_1("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,9,23,1,23,9,27,1,10,27,31,1,13,31,35,1,35,10,39,2,39,9,43,1,43,13,47,1,5,47,51,1,6,51,55,1,13,55,59,1,59,6,63,1,63,10,67,2,67,6,71,1,71,5,75,2,75,10,79,1,79,6,83,1,83,5,87,1,87,6,91,1,91,13,95,1,95,6,99,2,99,10,103,1,103,6,107,2,6,107,111,1,13,111,115,2,115,10,119,1,119,5,123,2,10,123,127,2,127,9,131,1,5,131,135,2,10,135,139,2,139,9,143,1,143,2,147,1,5,147,0,99,2,0,14,0")
	if result != "5534943" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "5534943")
	}
}
func TestDay2_2(t *testing.T) {
	result := Challenge2_2("1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,9,23,1,23,9,27,1,10,27,31,1,13,31,35,1,35,10,39,2,39,9,43,1,43,13,47,1,5,47,51,1,6,51,55,1,13,55,59,1,59,6,63,1,63,10,67,2,67,6,71,1,71,5,75,2,75,10,79,1,79,6,83,1,83,5,87,1,87,6,91,1,91,13,95,1,95,6,99,2,99,10,103,1,103,6,107,2,6,107,111,1,13,111,115,2,115,10,119,1,119,5,123,2,10,123,127,2,127,9,131,1,5,131,135,2,10,135,139,2,139,9,143,1,143,2,147,1,5,147,0,99,2,0,14,0")
	if result != "7603" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7603")
	}
}

func TestRunProgram_Example1(t *testing.T) {
	memory := CsvToIntSlice("1,0,0,0,99")
	RunProgram(&memory)
	msmCsv := GetMemoryAsCsv(memory)
	if msmCsv != "2,0,0,0,99" {
		t.Errorf("Memory was incorrect, got: %s, want: %s.", msmCsv, "2,0,0,0,99")
	}
}
func TestRunProgram_Example2(t *testing.T) {
	memory := CsvToIntSlice("2,3,0,3,99")
	RunProgram(&memory)
	msmCsv := GetMemoryAsCsv(memory)
	if msmCsv != "2,3,0,6,99" {
		t.Errorf("Memory was incorrect, got: %s, want: %s.", msmCsv, "2,3,0,6,99")
	}
}
func TestRunProgram_Example3(t *testing.T) {
	memory := CsvToIntSlice("2,4,4,5,99,0")
	RunProgram(&memory)
	msmCsv := GetMemoryAsCsv(memory)
	if msmCsv != "2,4,4,5,99,9801" {
		t.Errorf("Memory was incorrect, got: %s, want: %s.", msmCsv, "2,4,4,5,99,9801")
	}
}
func TestRunProgram_Example4(t *testing.T) {
	memory := CsvToIntSlice("1,1,1,4,99,5,6,0,99")
	RunProgram(&memory)
	msmCsv := GetMemoryAsCsv(memory)
	if msmCsv != "30,1,1,4,2,5,6,0,99" {
		t.Errorf("Memory was incorrect, got: %s, want: %s.", msmCsv, "30,1,1,4,2,5,6,0,99")
	}
}

func GetMemoryAsCsv(memory []int64) string {
	valuesText := []string{}

	for i := range memory {
		number := memory[i]
		text := strconv.FormatInt(number, 10)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, ",")
}
