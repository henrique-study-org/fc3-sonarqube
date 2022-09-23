package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(3, 2)

	if result != 1 {
		t.Error("The result must be 1")
	}
}
