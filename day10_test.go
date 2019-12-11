package main

import (
	"testing"
)

func TestDay10_1(t *testing.T) {
	input := ``
	output := Challenge10_1(input, 1)
	if output != "" {
		t.Errorf("Failure detected: %s", output)
	}
}
