package main

import (
	"testing"
)

func TestDay4_1(t *testing.T) {
	result := Challenge4_1("193651-649729")
	if result != "1605" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "1605")
	}
}

func TestExample4_1_1(t *testing.T) {
	if !MeetsRules("111111") {
		t.Errorf("Expected valid")
	}
}

func TestExample4_1_2(t *testing.T) {
	if MeetsRules("223450") {
		t.Errorf("Expected invalid")
	}
}

func TestExample4_1_3(t *testing.T) {
	if MeetsRules("123789") {
		t.Errorf("Expected invalid")
	}
}
