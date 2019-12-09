package main

import "testing"

func TestDay7_1(t *testing.T) {
	result := Challenge7_1("")
	if result != "402879" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "402879")
	}
}
