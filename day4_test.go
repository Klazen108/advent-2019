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

func TestDay4_2(t *testing.T) {
	result := Challenge4_2("193651-649729")
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

func TestExample4_2_1(t *testing.T) {
	if !MeetsRules2("112233") {
		t.Errorf("Expected valid")
	}
}

func TestExample4_2_2(t *testing.T) {
	if MeetsRules2("123444") {
		t.Errorf("Expected invalid")
	}
}

func TestExample4_2_3(t *testing.T) {
	if !MeetsRules2("111122") {
		t.Errorf("Expected valid")
	}
}

func TestExample4_2_x1(t *testing.T) {
	if !MeetsRules2("223333") {
		t.Errorf("Expected valid")
	}
}

func TestExample4_2_x2(t *testing.T) {
	if !MeetsRules2("344555") {
		t.Errorf("Expected valid")
	}
}

func TestExample4_2_x3(t *testing.T) {
	if MeetsRules2("123456") {
		t.Errorf("Expected invalid")
	}
}

func TestExample4_2_x4(t *testing.T) {
	if !MeetsRules2("113456") {
		t.Errorf("Expected valid")
	}
}
