package main

import "testing"

func TestDay12_1(t *testing.T) {
	output := Challenge12_1("")
	if output != "idk" {
		t.Errorf("Expected idk, got %s", output)
	}
}
