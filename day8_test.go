package main

import (
	"io/ioutil"
	"testing"
)

func TestDay8_1(t *testing.T) {
	input, _ := ioutil.ReadFile("challenge8-input.txt")
	result := Challenge8_1(input, 25, 6)
	if result != "2500" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "2500")
	}
}

func TestDay8_2(t *testing.T) {
	input, _ := ioutil.ReadFile("challenge8-input.txt")
	Challenge8_2(input, 25, 6)
	t.Errorf("Doing this so I can see the output")
}

func TestDay8_2_Example1(t *testing.T) {
	input := []byte("0222112222120000")
	Challenge8_2(input, 2, 2)
	t.Errorf("Doing this so I can see the output")
}
