package math

import "testing"

func TestAddInt(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddFloat(t *testing.T) {
	result := Add(2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("Add(2.5, 3.5) = %f; want %f", result, expected)
	}
}

func TestAddNegativeInt(t *testing.T) {
	result := Add(-2, -3)
	expected := -5
	if result != expected {
		t.Errorf("Add(-2, -3) = %d; want %d", result, expected)
	}
}

func TestAddNegativeFloat(t *testing.T) {
	result := Add(-2.5, -3.5)
	expected := -6.0
	if result != expected {
		t.Errorf("Add(-2.5, -3.5) = %f; want %f", result, expected)
	}
}

func TestAddMixedInt(t *testing.T) {
	result := Add(-2, 3)
	expected := 1
	if result != expected {
		t.Errorf("Add(-2, 3) = %d; want %d", result, expected)
	}
}

func TestAddMixedFloat(t *testing.T) {
	result := Add(-2.5, 3.5)
	expected := 1.0
	if result != expected {
		t.Errorf("Add(-2.5, 3.5) = %f; want %f", result, expected)
	}
}
