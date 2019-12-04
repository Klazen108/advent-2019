package main

import (
	"testing"
)

func TestDay3Example1(t *testing.T) {
	result := Challenge3_1("R8,U5,L5,D3\nU7,R6,D4,L4")
	if result != "6" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "6")
	}
}
