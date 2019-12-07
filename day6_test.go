package main

import (
	"testing"
)

func TestDay6_1(t *testing.T) {
	result := Challenge6_1("")
	if result != "" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "")
	}
}
