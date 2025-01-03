package math

import (
	"testing"
)

func TestMultiplyInt(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
	}
}

func TestMultiplyFloat64(t *testing.T) {
	result := Multiply(2.5, 4.0)
	expected := 10.0
	if result != expected {
		t.Errorf("Multiply(2.5, 4.0) = %f; want %f", result, expected)
	}
}
