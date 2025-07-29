package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		operator string
		expected string
	}{
		{"Addition", "0.1", "0.2", "+", "0.3"},
		{"Subtraction", "1.5", "2.5", "-", "-1.0"},
		{"Multiplication", "2", "3", "*", "6"},
		{"Division", "10", "4", "/", "2.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := decimal.NewFromString(tt.a)
			b, _ := decimal.NewFromString(tt.b)
			expected, _ := decimal.NewFromString(tt.expected)

			result, err := calc(a, tt.operator, b)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !result.Equal(expected) {
				t.Errorf("%s: %s %s %s = %s; want %s",
					tt.name, tt.a, tt.operator, tt.b, result.String(), expected.String())
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	t.Run("Divide by zero should return error", func(t *testing.T) {
		_, err := calc(decimal.NewFromInt(1), "/", decimal.NewFromInt(0))
		if err == nil {
			t.Errorf("Expected error, got nil")
		} else {
			expected := "cannot divide by zero"
			if err.Error() != expected {
				t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
			}
		}
	})
}

func TestInvalidOperator(t *testing.T) {
	t.Run("Invalid operator should return error", func(t *testing.T) {
		_, err := calc(decimal.NewFromInt(1), "%", decimal.NewFromInt(2))
		if err == nil {
			t.Errorf("Expected error for invalid operator, got nil")
		}
	})
}

func TestFloatingPointPrecision(t *testing.T) {
	t.Run("Floating point precision with decimal library", func(t *testing.T) {
		a, _ := decimal.NewFromString("0.10000000000000001")
		b, _ := decimal.NewFromString("0.20000000000000001")
		expected, _ := decimal.NewFromString("0.30000000000000002")

		result, err := calc(a, "+", b)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
