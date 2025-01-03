package math

import (
	"testing"
)

func TestSubtractInt(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{10, 5, 5},
		{0, 0, 0},
		{-5, -5, 0},
		{100, 50, 50},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtractFloat64(t *testing.T) {
	tests := []struct {
		a, b, expected float64
	}{
		{10.5, 5.5, 5.0},
		{0.0, 0.0, 0.0},
		{-5.5, -5.5, 0.0},
		{100.0, 50.0, 50.0},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%f, %f) = %f; want %f", test.a, test.b, result, test.expected)
		}
	}
}
