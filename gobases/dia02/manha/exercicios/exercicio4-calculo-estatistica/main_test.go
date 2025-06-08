package main

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.00001
}

func TestMinValue(t *testing.T) {
	type testCase struct {
		input    []float64
		expected float64
	}
	tests := []testCase{
		{[]float64{2, 3, 1, 5, 4}, 1},
		{[]float64{10, 2, 10, 10}, 2},
		{[]float64{5, 5, 5, 5}, 5},
	}
	for _, tc := range tests {
		got := minValue(tc.input...)
		if !floatEquals(got, tc.expected) {
			t.Errorf("minValue(%v) = %v; esperado %v", tc.input, got, tc.expected)
		}
	}
}

func TestMaxValue(t *testing.T) {
	type testCase struct {
		input    []float64
		expected float64
	}
	tests := []testCase{
		{[]float64{2, 3, 1, 5, 4}, 5},
		{[]float64{10, 2, 10, 10}, 10},
		{[]float64{5, 5, 5, 5}, 5},
	}
	for _, tc := range tests {
		got := maxValue(tc.input...)
		if !floatEquals(got, tc.expected) {
			t.Errorf("maxValue(%v) = %v; esperado %v", tc.input, got, tc.expected)
		}
	}
}

func TestAvgValue(t *testing.T) {
	type testCase struct {
		input    []float64
		expected float64
	}
	tests := []testCase{
		{[]float64{2, 3, 1, 5, 4}, 3},
		{[]float64{10, 0}, 5},
		{[]float64{5, 5, 5, 5}, 5},
	}
	for _, tc := range tests {
		got := avgValue(tc.input...)
		if !floatEquals(got, tc.expected) {
			t.Errorf("avgValue(%v) = %v; esperado %v", tc.input, got, tc.expected)
		}
	}
}

func TestOperationFunc(t *testing.T) {
	minFunc, err := operation(minimum)
	if err != nil {
		t.Fatalf("Esperava nil error para minimum, obteve %v", err)
	}
	if minFunc(2, 1, 9) != 1 {
		t.Errorf("operation(minimum) não retornou função correta")
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		t.Fatalf("Esperava nil error para maximum, obteve %v", err)
	}
	if maxFunc(2, 1, 9) != 9 {
		t.Errorf("operation(maximum) não retornou função correta")
	}

	avgFunc, err := operation(average)
	if err != nil {
		t.Fatalf("Esperava nil error para average, obteve %v", err)
	}
	result := avgFunc(2, 4, 6)
	if !floatEquals(result, 4.0) {
		t.Errorf("operation(average) não retornou função correta. Esperado 4.0, obteve %v", result)
	}
}

func TestOperationInvalid(t *testing.T) {
	_, err := operation("sum")
	if err == nil {
		t.Errorf("Esperava erro para operação inválida, obteve nil")
	}
}
