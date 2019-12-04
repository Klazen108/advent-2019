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

func TestDay3Example2(t *testing.T) {
	result := Challenge3_1("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	if result != "159" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "159")
	}
}
