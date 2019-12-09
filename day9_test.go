package main

import (
	"testing"
)

func TestDay9_1(t *testing.T) {
	result := Challenge9_1("")
	if result != "2500" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "2500")
	}
}
