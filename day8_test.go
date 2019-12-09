package main

import (
	"io/ioutil"
	"testing"
)

func TestDay8_1(t *testing.T) {
	input, _ := ioutil.ReadFile("challenge8-input.txt")
	result := Challenge8_1(input)
	if result != "2500" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "2500")
	}
}
