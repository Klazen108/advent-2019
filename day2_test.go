package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestChallenge2_1Example(t *testing.T) {
	memory := Challenge2_1("1,9,10,3,2,3,11,0,99,30,40,50")

	valuesText := []string{}

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range memory {
		number := memory[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	result := strings.Join(valuesText, ",")

	if result != "30,1,1,4,2,5,6,0,99" {
		result := strings.Join(valuesText, ",")
		t.Errorf("Memory was incorrect, got: %s, want: %s.", result, "30,1,1,4,2,5,6,0,99")
	}
}
