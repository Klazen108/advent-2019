package main

import (
	"testing"
)

func TestDay4_1(t *testing.T) {
	result := Challenge4_1("193651-649729")
	if result != "" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "")
	}
}
