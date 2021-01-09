package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculateCorrect(t *testing.T) {
	if Calculate(3) != 4 {
		t.Error("Expected value 4: Actual value 5")
	}
}
