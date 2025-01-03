package math

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		method   string
		a, b     float64
		expected float64
	}{
		{"sum", 1, 2, 3},
		{"sub", 5, 3, 2},
		{"div", 6, 2, 3},
		{"mult", 2, 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			result := Calculator(tt.method, tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Calculator(%s, %f, %f) = %f; want %f", tt.method, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculatorInvalidMethod(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid method")
		}
	}()
	Calculator("invalid", 1, 2)
}
