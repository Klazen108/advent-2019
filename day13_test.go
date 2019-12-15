package main

import "testing"

func TestDay13_1(t *testing.T) {
	output := Challenge13_1()
	if output != "286" {
		t.Errorf("Expected 286, got %s", output)
	}
}
func TestDay13_2(t *testing.T) {
	output := Challenge13_2()
	if output != "286" {
		t.Errorf("Expected 286, got %s", output)
	}
}
