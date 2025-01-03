package math

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     interface{}
		expected interface{}
		wantErr  bool
	}{
		{"int division", 10, 2, 5, false},
		{"float division", 10.0, 2.0, 5.0, false},
		{"division by zero", 10, 0, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Divide() panicked unexpectedly: %v", r)
					}
				}
			}()

			var result interface{}
			switch a := tt.a.(type) {
			case int:
				result = Divide(a, tt.b.(int))
			case float64:
				result = Divide(a, tt.b.(float64))
			}

			if !tt.wantErr && result != tt.expected {
				t.Errorf("Divide() = %v, want %v", result, tt.expected)
			}
		})
	}
}
